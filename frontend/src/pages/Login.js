import React, { useState } from 'react';
import api from '../services/api';
import { useNavigate } from 'react-router-dom';

const Login = () => {
	const [email, setEmail] = useState('');
	const [password, setPassword] = useState('');
	const [message, setMessage] = useState('');
	const navigate = useNavigate();

	const handleLogin = async () => {
		try {
			const response = await api.post('/auth/login', { email, password });
			setMessage('Login successful');
			navigate('/');
		} catch (error) {
			if (error.response && error.response.status === 401) {
				setMessage('Invalid credentials');
			} else {
				setMessage('Login failed');
			}
		}
	};

	return (
		<div>
			<h2>Login</h2>
			<input
				type="text"
				placeholder="Email"
				value={email}
				onChange={(e) => setEmail(e.target.value)}
			/>
			<input
				type="password"
				placeholder="Password"
				value={password}
				onChange={(e) => setPassword(e.target.value)}
			/>
			<button onClick={handleLogin}>Login</button>
			<p>{message}</p>
		</div>
	);
};

export default Login;
