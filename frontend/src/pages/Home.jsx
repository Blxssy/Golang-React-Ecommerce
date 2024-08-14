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
						<p className="product-name">Product 1</p>
						<p className="product-description">Description 1</p>
					</div>
					<div className='product-item'>
						<p className="product-name">Product 2</p>
						<p className="product-description">Description 2</p>
					</div>
					<div className='product-item'>
						<p className="product-name">Product 3</p>
						<p className="product-description">Description 3</p>
					</div>
					<div className='product-item'>
						<p className="product-name">Product 4</p>
						<p className="product-description">Description 4</p>
					</div>
					<div className='product-item'>
						<p className="product-name">Product 5</p>
						<p className="product-description">Description 5</p>
					</div>
					<div className='product-item'>
						<p className="product-name">Product 6</p>
						<p className="product-description">Description 6</p>
					</div>
					<div className='product-item'>
						<p className="product-name">Product 7</p>
						<p className="product-description">Description 7</p>
					</div>
					<div className='product-item'>
						<p className="product-name">Product 8</p>
						<p className="product-description">Description 8</p>
					</div>
					<div className='product-item'>
						<p className="product-name">Product 9</p>
						<p className="product-description">Description 9</p>
					</div>
					<div className='product-item'>
						<p className="product-name">Product 10</p>
						<p className="product-description">Description 10</p>
					</div>
					<div className='product-item'>
						<p className="product-name">Product 11</p>
						<p className="product-description">Description 11</p>
					</div>
					<div className='product-item'>
						<p className="product-name">Product 12</p>
						<p className="product-description">Description 12</p>
					</div>
				</div>
			</section>
		</div>
	);
};


export default Home;