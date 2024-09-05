// import React, { useEffect, useState } from 'react';
// import "../styles/Cart.css";

// const Cart = () => {
//   const [cartItems, setCartItems] = useState([]);

//   useEffect(() => {
//     fetch('http://localhost:3001/api/cart')
//       .then(response => response.json())
//       .then(data => {
//         if (Array.isArray(data)) {
//           setCartItems(data);
//         } else {
//           console.error('Expected an array but got:', data);
//         }
//       })
//       .catch(error => console.error('Error fetching cart items:', error));
//   }, []);

//   return (
//     <div className="cart-container"> 
//       <h1>Your Cart</h1>
//       <ul>
//         {cartItems.map(item => (
//           <li key={item.product_id}>
//             Product ID: {item.product_id}, Quantity: {item.quantity}
//           </li>
//         ))}
//       </ul>
//     </div>
//   );
// };

// export default Cart;


import React, { useEffect, useState } from 'react';
import "../styles/Cart.css";

const Cart = () => {
  const [cartItems, setCartItems] = useState([]);

  useEffect(() => {
    const fetchCartItems = async () => {
      try {
        const accessToken = getCookie('access_token'); 

        const response = await fetch('http://localhost:3001/api/cart', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'Authorization': `Bearer ${accessToken}` 
          },
          credentials: 'include' 
        });

        const data = await response.json();
        if (data && data.userCart && Array.isArray(data.userCart.items)) {
          const extractedItems = data.userCart.items.map(item => ({
            product_id: item.product_id,
            quantity: item.quantity
          }));
          setCartItems(extractedItems);
        } else {
          console.error('Unexpected data format:', data);
        }
      } catch (error) {
        console.error('Error fetching cart items:', error);
      }
    };

    fetchCartItems();
  }, []);

  return (
    <div className="cart-container"> 
      <h1>Your Cart</h1>
      <ul>
        {cartItems.map(item => (
          <li key={item.product_id}>
            Product ID: {item.product_id}, Quantity: {item.quantity}
          </li>
        ))}
      </ul>
    </div>
  );
};

const getCookie = (name) => {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(';').shift();
};

export default Cart;
