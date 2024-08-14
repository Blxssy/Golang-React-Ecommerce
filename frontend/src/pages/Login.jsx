import React, { useState } from 'react';
import api from '../services/api';
import { useNavigate } from 'react-router-dom';
import './Login.css'; 

const Login = () => {
	const [email, setEmail] = useState('');
	const [password, setPassword] = useState('');
	const [message, setMessage] = useState('');
	const navigate = useNavigate();

	const handleLogin = async () => {
		try {
			await api.post('/auth/login', { email, password });
			setMessage('Login successful');
			navigate('/');
			window.location.reload();
		} catch (error) {
			if (error.response && error.response.status === 401) {
				setMessage('Invalid credentials');
			} else {
				setMessage('Login failed, try another email or password');
			}
		}
	};

	return (
        <div className="login">
            <div className="banner">
                <h1>Welcome back</h1>
                <p>Log in to proceed</p>
            </div>
            <form onSubmit={(e) => { e.preventDefault();}}>
                <div className="input-group">
                    <label htmlFor="email"> </label>
                    <input
                        id="email"
                        type="text"
                        placeholder="Enter your email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                </div>
                <div className="input-group">
                    <label htmlFor="password"> </label>
                    <input
                        id="password"
                        type="password"
                        placeholder="Enter your password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </div>
                <button onClick={handleLogin}>Login</button>
                <p>{message}</p>
            </form> 
        </div> 
    );

	// return (
	// 	<div>
	// 		<h2>Login</h2>
	// 		<input
	// 			type="text"
	// 			placeholder="Email"
	// 			value={email}
	// 			onChange={(e) => setEmail(e.target.value)}
	// 		/>
	// 		<input
	// 			type="password"
	// 			placeholder="Password"
	// 			value={password}
	// 			onChange={(e) => setPassword(e.target.value)}
	// 		/>
	// 		<button onClick={handleLogin}>Login</button>
	// 		<p>{message}</p>
	// 	</div>
	// );
};

export default Login;