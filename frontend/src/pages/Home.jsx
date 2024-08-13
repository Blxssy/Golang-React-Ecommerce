import React from 'react';
import './Home.css'

const Home = () => {
	return (
		<div className='home'>
			<div className='banner'>
				<h1>Welcome to Rom4ik's Store</h1>
				<p>Discover amazing products we created just for you</p>
			</div>
			<section className='products'>
				<h2>Products List</h2>
				<div className='product-grid'>
					<div className='product-item'>
						Product 1
						<p className="product-description">Description 1</p>
					</div>
					<div className='product-item'>
						Product 2
						<p className="product-description">Description 2</p>
					</div>
					<div className='product-item'>
						Product 3
						<p className="product-description">Description 3</p>
					</div>
					<div className='product-item'>
						Product 4
						<p className="product-description">Description 4</p>
					</div>
				</div>
			</section>
		</div>
	);
};


export default Home;