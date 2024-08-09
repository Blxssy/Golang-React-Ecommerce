// src/pages/Home.js
import React, {useState} from 'react'
import { useNavigate } from 'react-router-dom'
import { ProductsList } from '../components/ProductsList'
import './Home.css'
import api from "../services/api";

const Home = () => {
	const [email, setEmail] = useState('');
	const [password, setPassword] = useState('');
	const [message, setMessage] = useState('');

	const handleRegister = async () => {
		try {
			const response = await api.post('/auth/register', { email, password });
			setMessage(response.data.message);
		} catch (error) {
			setMessage('Registration failed');
		}
	};

	const handleLogin = async () => {
		try {
			const response = await api.post('/auth/login', { email, password });
			setMessage(response.data.message);
		} catch (error) {
			if (error.response && error.response.status === 401) {
				setMessage('Invalid credentials');
			} else {
				setMessage('Login failed');
			}
		}
	};

	const handleProtectedRequest = async () => {
		try {
			const response = await api.get('/protected-route');
			setMessage(response.data.message);
		} catch (error) {
			setMessage('Failed to fetch protected data');
		}
	};

	return (
		<div className='home'>
			<h1>Products List</h1>
			<div>
				<input type="text" placeholder="email" onChange={(e) => setEmail(e.target.value)}/>
				<input type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)}/>
				<button onClick={handleRegister}>Register</button>
				<button onClick={handleLogin}>Login</button>
				<button onClick={handleProtectedRequest}>Fetch Protected Data</button>
				<p>{message}</p>
			</div>
			<ProductsList/>
		</div>
	)
}

export default Home
