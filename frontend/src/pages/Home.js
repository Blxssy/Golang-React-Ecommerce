import React from 'react';
import { ProductsList } from '../components/ProductsList';
import '../styles/Home.css';

const Home = () => {
	return (
		<div className='home'>
			<h1>Products List</h1>
			<ProductsList/>
		</div>
	);
};

export default Home;
