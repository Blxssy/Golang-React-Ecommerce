import React from 'react';
import '../styles/Home.css';

const Home = () => {
	return (
		<div className='home'>
			<div className='banner'>
				<h1>Welcome to Our Store</h1>
				<p>Discover amazing products we created just for you</p>
			</div>
			<section className='products'>
				<h2>Products List</h2>
				<div className='product-grid'>
					{/* Product items will go here */}
					<div className='product-item'>Product 1</div>
					<div className='product-item'>Product 2</div>
					<div className='product-item'>Product 3</div>
					<div className='product-item'>Product 4</div>
				</div>
			</section>
		</div>
	);
};

export default Home;