import React, { useState, useEffect } from 'react';
import '../styles/Home.css';
// import '../styles/Сarousel.css';
import '../styles/ProductCard.css'
import api from "../services/api";

const Home = () => {
  const [user, setUser] = useState(null);
  const [products, setProducts] = useState([]);
  const [currentIndex, setCurrentIndex] = useState(0);
  const [touchStart, setTouchStart] = useState(null);
  const [touchMove, setTouchMove] = useState(null);

  // useEffect(() => {
  //   const intervalId = setInterval(() => {
  //     setCurrentIndex((prevIndex) => (prevIndex + 1) % products.length);
  //   }, 3000); 
  //   return () => clearInterval(intervalId);
  // }, [products]);

  const handlePrevClick = () => {
    setCurrentIndex((prevIndex) => (prevIndex - 1 + products.length) % products.length);
  };

  const handleNextClick = () => {
    setCurrentIndex((prevIndex) => (prevIndex + 1) % products.length);
  };

  const handleTouchStart = (event) => {
    setTouchStart(event.touches[0].clientX);
  };

  const handleTouchMove = (event) => {
    setTouchMove(event.touches[0].clientX);
  };

  const handleTouchEnd = () => {
    if (touchStart && touchMove) {
      const diff = touchMove - touchStart;
      if (diff > 50) {
        handlePrevClick();
      } else if (diff < -50) {
        handleNextClick();
      }
    }
    setTouchStart(null);
    setTouchMove(null);
  };

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

  const repeatedProducts = [...products, ...products];

  return (
    <div className='home'>
		<div className="greeting-container">
		<h1 className="h1">
			Welcome to our store, {user ? user.username : "Guest"}!
		</h1>
    <button className="prev" onClick={handlePrevClick}>
      Prev
    </button>
    <button className="next" onClick={handleNextClick}>
      Next
    </button>
    <div className="carousel-container">
      <div
        className="carousel"
        onTouchStart={handleTouchStart}
        onTouchMove={handleTouchMove}
        onTouchEnd={handleTouchEnd}
      >
        {repeatedProducts.map((product, index) => (
          <div
            key={index}
            className="product-item"
            style={{
              transform: `translateX(${(index - currentIndex) * 100}%)`,
              transition: 'transform 0.3s ease-in-out',
            }}
          >
            <h3>{product.name}</h3>
            <img src={product.image} alt={product.name} />
          </div>
        ))}
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
