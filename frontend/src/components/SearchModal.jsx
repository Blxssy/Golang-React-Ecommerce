import React, { useState, useEffect } from 'react';
import './SearchModal.css';

const SearchModal = ({ isOpen, onClose }) => {
  const [query, setQuery] = useState('');
  const [searchResults, setSearchResults] = useState([]);
  const [history, setHistory] = useState([]);

  useEffect(() => {
    const savedHistory = JSON.parse(localStorage.getItem('searchHistory')) || [];
    setHistory(savedHistory);
  }, []);

  const handleSearch = async (e) => {
    e.preventDefault();
    if (!query.trim()) return;

    const newHistory = [query, ...history.filter((item) => item !== query)];
    setHistory(newHistory);
    localStorage.setItem('searchHistory', JSON.stringify(newHistory));

    try {
      const response = await fetch(`http://localhost:3001/api/products?name=${query}`);
      if (response.ok) {
        const data = await response.json();
        setSearchResults(data); 
      } else {
        console.error('Ошибка при запросе к API');
      }
    } catch (error) {
      console.error('Ошибка:', error);
    }

    setQuery('');
  };

  return (
    isOpen && (
      <div className="search-modal">
        <div className="search-modal-overlay" onClick={onClose} />
        <div className="search-modal-content">
          <form onSubmit={handleSearch}>
            <input
              type="text"
              placeholder="Search..."
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              className="search-input"
            />
            <button type="submit">Search</button>
          </form>

          {searchResults.length > 0 && (
            <div className="search-results">
              <h3>Search Results:</h3>
              <ul>
                {searchResults.map((item) => (
                  <li key={item.id}>{item.name}</li>
                ))}
              </ul>
            </div>
          )}

          {history.length > 0 && (
            <div className="search-history">
              <h3>Search History:</h3>
              <ul>
                {history.map((item, index) => (
                  <li key={index}>{item}</li>
                ))}
              </ul>
            </div>
          )}
        </div>
      </div>
    )
  );
};

export default SearchModal;
