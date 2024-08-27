// import React from 'react';
// import { useLocation } from 'react-router-dom';

// const ProductDetailPage = () => {
//   const location = useLocation();
//   const { product } = location.state || {}; 

//   if (!product) {
//     return <div>Product not found</div>;
//   }

//   return (
//     <div>
//       <h1>{product.name}</h1>
//       <img src={product.image} alt={product.name} />
//       <p>{product.description}</p>
//       <p>Price: {product.price} $</p>
//     </div>
//   );
// };

// export default ProductDetailPage;

import React from 'react';
import { useLocation } from 'react-router-dom';
import '/Users/Amogus/Golang-React-Ecommerce/frontend/src/styles/ProductDetailPage.css';

const ProductDetailPage = () => {
  const location = useLocation();
  const { product } = location.state || {}; 

  if (!product) {
    return <div>Product not found</div>;
  }

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
        <button className="add-to-cart-button">Add to Cart</button>
      </div>
    </div>
  );
};

export default ProductDetailPage;

