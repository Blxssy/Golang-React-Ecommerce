import React from 'react';
import '../styles/BottomMenu.css';
import {useNavigate} from "react-router-dom";
import { FaThLarge, FaSearch, FaShoppingCart } from 'react-icons/fa'; // Используем библиотеку react-icons
import { IoIosHome } from "react-icons/io";

const BottomMenu = () => {
  const navigate = useNavigate();

  return (
    <div className="menu-container">
      <div className="menu-item" onClick={() => navigate('/')}>
        <IoIosHome size={30} />
      </div>
      <div className="menu-item" onClick={() => navigate('/products')}>
        <FaThLarge size={30} />
      </div>
      <div className="menu-item" onClick={() => navigate('/search')}>
        <FaSearch size={30} />
      </div>
      <div className="menu-item" onClick={() => navigate('/cart')}>
        <FaShoppingCart size={30} />
      </div>
    </div>
  );
};

export default BottomMenu;
