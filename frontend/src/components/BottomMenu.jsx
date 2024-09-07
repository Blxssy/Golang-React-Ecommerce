import React from 'react';
import '../styles/BottomMenu.css';
import { IoIosHome, FaThLarge, FaSearch, FaShoppingCart } from 'react-icons/fa'; // Используем библиотеку react-icons

const BottomMenu = () => {
  return (
    <div className="menu-container">
      <div className="menu-item">
        <IoIosHome size={30} />
      </div>
      <div className="menu-item">
        <FaThLarge size={30} />
      </div>
      <div className="menu-item">
        <FaSearch size={30} />
      </div>
      <div className="menu-item">
        <FaShoppingCart size={30} /> 
      </div>
    </div>
  );
};

export default BottomMenu;
