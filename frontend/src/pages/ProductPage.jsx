// import React from 'react';
// import { useLocation } from 'react-router-dom';
// import '../styles/ProductDetailPage.css';

// const ProductDetailPage = () => {
//   const location = useLocation();
//   const { product } = location.state || {}; 

//   if (!product) {
//     return <div>Product not found</div>;
//   }

//   return (
//     <div className="product-detail-container">
//       <div className="product-image-section">
//         <img className="product-image" src={product.image} alt={product.name} />
//       </div>
//       <div className="product-info-section">
//         <h1 className="product-name">{product.name}</h1>
//         <p className="product-price">Price: <strong>{product.price} $</strong></p>
//         <p className="product-description">{product.description}</p>
//         <button className="buy-now-button">Buy Now</button>
//         <button className="add-to-cart-button">Add to Cart</button>
//       </div>
//     </div>
//   );
// };

// export default ProductDetailPage;

import React from 'react';
import { useLocation } from 'react-router-dom';
import '../styles/ProductDetailPage.css';

const ProductDetailPage = () => {
  const location = useLocation();
  const { product } = location.state || {}; 

  if (!product) {
    return <div>Product not found</div>;
  }

  const handleAddToCart = async () => {
    // try {
    //   const accessToken = getCookie('access_token'); 

    //   const response = await fetch('http://localhost:3001/api/cart/items', {
    //     method: 'POST',
    //     headers: {
    //       'Content-Type': 'application/json',
    //       'Accept': 'application/json',
    //       'Authorization': `Bearer ${accessToken}`
    //     },
    //     body: JSON.stringify({ 
    //       product_id: product.id,  
    //       quantity: 1
    //     }), 
    //   });

    //   if (response.ok) {
    //     console.log('Product added to cart successfully!');
    //   } else {
    //     console.error('Failed to add product to cart');
    //   }
    // } catch (error) {
    //   console.error('Error adding product to cart:', error);
    // }
    try {
      const accessToken = getCookie('access_token'); 
  
      const response = await fetch('http://localhost:3001/api/cart/items', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'accept': 'application/json',
          'access_token': accessToken, 
          // 'Authorization': `Bearer ${accessToken}`
        },
        body: JSON.stringify({ 
          product_id: product.id,  
          quantity: 1  
        }), 
      });
  
      if (response.ok) {
        console.log('Product added to cart successfully!');
      } else {
        console.error('Failed to add product to cart', response.status, response.statusText);
      }
    } catch (error) {
      console.error('Error adding product to cart:', error);
    }
  };


  return (
    <div className="product-detail-container">
      <div className="product-image-section">
        <img className="product-image" src={product.image} alt={product.name} />
      </div>
      <div className="product-info-section">
        <h1 className="product-name">{product.name}</h1>
        <p className="product-price">Price: <strong>{product.price} $</strong></p>
        <p className="product-description">{product.description}</p>
        <button className="buy-now-button">Buy Now</button>
        <button className="add-to-cart-button" onClick={handleAddToCart}>
          Add to Cart
        </button>
      </div>
    </div>
  );
};

const getCookie = (name) => {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(';').shift();
};

export default ProductDetailPage;
