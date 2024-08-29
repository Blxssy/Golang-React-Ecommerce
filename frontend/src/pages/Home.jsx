import React, { useState, useEffect } from 'react';
import '../styles/Home.css';
import '../styles/Сarousel.css'
import api from "../services/api";

const Home = () => {
  const [user, setUser] = useState(null);
  const [products, setProducts] = useState([]);

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await api.get('/auth/user-info');
        setUser(response.data);
      } catch (error) {
        console.log('Not authenticated');
      }
    };
    fetchUser(); 
  }, []);
  
  useEffect(() => {
    fetch('http://localhost:3001/api/products')
      .then(response => response.json())
      .then(data => setProducts(data))
      .catch(error => console.error('Error loading products:', error));
  }, []); 

  const scrollToMiddle = () => {
    const middle = document.documentElement.scrollHeight / 2;
    window.scrollTo({
      top: middle,
      behavior: 'smooth'
    });
  };

  const scrollToBottom = () => {
    window.scrollTo({
      top: document.documentElement.scrollHeight,
      behavior: 'smooth'
    });
  };

  return (
    <div className='home'>
		<div className="greeting-container">
		<h1 className="h1">
			Welcome to our store, {user ? user.username : "Guest"}!
		</h1>
    <div className="carousel-container">
      <div className="carousel">
        {products.length > 0 ? (
          products.map((product, index) => (
            <div key={index} className="carousel-item">
              <h3>{product.name}</h3>
              <img
                src={product.image} 
                alt={product.name} 
              />
            </div>
          ))
        ) : (
          <p>Loading products...</p>
        )}
      </div>
    </div>
		<button className="scroll-button" onClick={scrollToMiddle}>
			Want more? ↓
		</button>
		</div>
		
		<div className="middle-button-container">
  			<h1>Discover what we've got just for you!</h1>
  			<button className="scroll-button" onClick={scrollToBottom}>
    			Intrested? ↓
  			</button>
		</div>

      <div className="bottom-button-container">
        <a href="/products" className="bebr-button">Dive in</a>
      </div>
    </div>
  );
};

export default Home;
