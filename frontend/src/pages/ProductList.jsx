import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
// import '/Users/Amogus/Golang-React-Ecommerce/frontend/src/styles/Home.css'
import '/Users/Amogus/Golang-React-Ecommerce/frontend/src/styles/ProductCard.css'

const ProductList = () => {
  const [products, setProducts] = useState([]);
  const navigate = useNavigate(); 

  useEffect(() => {
    fetch('http://localhost:3001/api/products')
      .then(response => response.json())
      .then(data => setProducts(data))
      .catch(error => console.error('Error loading products:', error));
  }, []);

  const handleProductClick = (product) => {
    navigate(`/products/${product.id}`, { state: { product } });
  };

  return (
    <div className="products-container">
      <div className="product-grid">
        {products.map(product => (
            <div className="product-item"
            onClick={() => handleProductClick(product)}
            >
              <img 
                src={product.image} 
                alt={product.name} 
              />
              <h3>{product.name}</h3>
              <p>Price: {product.price} $</p>
            </div>
        ))}
      </div>
    </div>
  );
};

export default ProductList;
