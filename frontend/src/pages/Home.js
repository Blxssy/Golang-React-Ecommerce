// src/pages/Home.js
import React from 'react'
import { useNavigate } from 'react-router-dom'
import { ProductsList } from '../components/ProductsList'
import './Home.css'

const Home = () => {
	const navigate = useNavigate()

	const handleRegister = () => {
		navigate('/register')
	}

	const handleLogin = () => {
		navigate('/login')
	}

	return (
		<div className='home'>
			<h1>Products List</h1>
			<div className='auth-buttons'>
				<button onClick={handleRegister}>Register</button>
				<button onClick={handleLogin}>Login</button>
			</div>
			<ProductsList />
		</div>
	)
}

export default Home
