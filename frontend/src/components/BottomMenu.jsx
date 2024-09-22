// import React, { useState } from 'react';
import '../styles/BottomMenu.css';
import {useNavigate} from "react-router-dom";
import { FaThLarge, FaSearch, FaShoppingCart } from 'react-icons/fa'; 
import { IoIosHome } from "react-icons/io";
// import SearchModal from './SearchModal';

const BottomMenu = () => {
  const navigate = useNavigate();
  // const [isSearchOpen, setIsSearchOpen] = useState(false);

  return (
    <div className="menu-container">
      <div className="menu-item" onClick={() => navigate('/')}>
        <IoIosHome size={30} />
      </div>
      <div className="menu-item" onClick={() => navigate('/products')}>
        <FaThLarge size={30} />
      </div>
      {/* <div className="menu-item" onClick={() => setIsSearchOpen(true)}>
        <FaSearch size={30} />
      </div> */}
      <div className="menu-item" onClick={() => navigate('/search')}>
        <FaSearch size={30} />
      </div>
      <div className="menu-item" onClick={() => navigate('/cart')}>
        <FaShoppingCart size={30} />
      </div>
      {/* <SearchModal isOpen={isSearchOpen} onClose={() => setIsSearchOpen(false)} /> */}
    </div>
  );
};

export default BottomMenu;
