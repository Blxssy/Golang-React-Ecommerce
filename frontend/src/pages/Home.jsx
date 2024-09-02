// страница с кривым слайдером снизу
// import React, { useState, useEffect, useRef } from 'react';
// import '../styles/Home.css';
// import '../styles/ProductCard.css';
// import api from "../services/api";

// const Home = () => {
//   const [user, setUser] = useState(null);
//   const [products, setProducts] = useState([]);
//   const [currentIndex, setCurrentIndex] = useState(1); // Start with the second element (first clone)
//   const [isTransitioning, setIsTransitioning] = useState(false);
//   const carouselRef = useRef(null);

//   const step = 0.5;

//   useEffect(() => {
//     const fetchUser = async () => {
//       try {
//         const response = await api.get('/auth/user-info');
//         setUser(response.data);
//       } catch (error) {
//         console.log('Not authenticated');
//       }
//     };
//     fetchUser(); 
//   }, []);
  
//   useEffect(() => {
//     fetch('http://localhost:3001/api/products')
//       .then(response => response.json())
//       .then(data => {
//         // Clone the first and last product for the infinite loop effect
//         setProducts([data[data.length - 1], ...data, data[0]]);
//       })
//       .catch(error => console.error('Error loading products:', error));
//   }, []); 

//   const scrollToMiddle = () => {
//     const middle = document.documentElement.scrollHeight / 2;
//       window.scrollTo({
//         top: middle,
//         behavior: 'smooth'
//         });
//       };
    
//   const scrollToBottom = () => {
//     window.scrollTo({
//       top: document.documentElement.scrollHeight,
//       behavior: 'smooth'
//       });
//     };

//   const handlePrevClick = () => {
//     if (isTransitioning) return;
//     setIsTransitioning(true);
//     setCurrentIndex(prevIndex => prevIndex - step);
//   };

//   const handleNextClick = () => {
//     if (isTransitioning) return;
//     setIsTransitioning(true);
//     setCurrentIndex(prevIndex => prevIndex + step);
//   };

//   const handleTransitionEnd = () => {
//     setIsTransitioning(false);
//     // Handle the infinite loop by moving the index without transition
//     if (currentIndex === 0) {
//       setCurrentIndex(products.length - 1);
//       carouselRef.current.style.transition = 'none';
//       carouselRef.current.style.transform = `translateX(-${(products.length - 1) * 100}%)`;
//     } else if (currentIndex === products.length - 1) {
//       setCurrentIndex(1);
//       carouselRef.current.style.transition = 'none';
//       carouselRef.current.style.transform = `translateX(-100%)`;
//     }
//   };

//   useEffect(() => {
//     if (!isTransitioning) {
//       carouselRef.current.style.transition = 'transform 0.3s ease-in-out';
//     }
//     carouselRef.current.style.transform = `translateX(-${currentIndex * 100}%)`;
//   }, [currentIndex, isTransitioning]);

//   return (
//     <div className='home'>
// 		<div className="greeting-container">
// 		<h1 className="h1">
// 			Welcome to our store, {user ? user.username : "Guest"}!
// 		</h1>
//     <button className="prev" onClick={handlePrevClick}>
//       Prev
//     </button>
//     <button className="next" onClick={handleNextClick}>
//       Next
//     </button>
//     <div className="carousel-container">
//       <div
//         className="carousel"
//         ref={carouselRef}
//         onTransitionEnd={handleTransitionEnd}
//       >
//         {products.map((product, index) => {
//           const isActive = index === currentIndex;
//           return (
//             <div
//               key={index}
//               className={`product-item ${isActive ? 'active' : ''}`}
//               style={{
//                 transform: `translateX(${(index - currentIndex) * 10}%)`,
//                 transition: 'transform 0.3s ease-in-out',
//               }}
//             >
//               <h3>{product.name}</h3>
//               <img src={product.image} alt={product.name} />
//             </div>
//           );
//         })}
//       </div>
//     </div>
// 		<button className="scroll-button" onClick={scrollToMiddle}>
// 			Want more? ↓
// 		</button>
// 		</div>
		
// 		<div className="middle-button-container">
//   			<h1>Discover what we've got just for you!</h1>
//   			<button className="scroll-button" onClick={scrollToBottom}>
//     			Interested? ↓
//   			</button>
// 		</div>

//       <div className="bottom-button-container">
//         <a href="/products" className="bebr-button">Dive in</a>
//       </div>
//     </div>
//   );
// };

// export default Home;


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