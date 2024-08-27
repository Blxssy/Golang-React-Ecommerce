import React, { useState, useEffect } from 'react';
import '/Users/Amogus/Golang-React-Ecommerce/frontend/src/styles/Home.css';
import api from "../services/api";

const Home = () => {
  const [user, setUser] = useState(null);

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
		<button className="scroll-button" onClick={scrollToMiddle}>
			What's next? ↓
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
