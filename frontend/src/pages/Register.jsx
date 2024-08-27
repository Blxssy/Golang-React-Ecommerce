import React, { useState } from 'react';
import api from '../services/api';
import { useNavigate } from 'react-router-dom';
import '/Users/Amogus/Golang-React-Ecommerce/frontend/src/styles/Register.css';
import Cookies from "js-cookie";

const Register = () => {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [message, setMessage] = useState('');
    const navigate = useNavigate();

    const handleRegister = async () => {
        try {
            const response = await api.post('/auth/register', { username, email, password });
            const { access_token, refresh_token } = response.data;

            Cookies.set('access_token', access_token);
            Cookies.set('refresh_token', refresh_token);

            setMessage('Register successful');

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
        <div className="register">
            <div className="banner">
                <h1>Registration</h1>
                <p>Join to discover more</p>
            </div>
            <form onSubmit={(e) => { e.preventDefault(); handleRegister(); }}>
            <div className="input-group">
                    <label htmlFor="username"> </label>
                    <input
                        id="username"
                        type="text"
                        placeholder="Enter your username"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                </div>
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
                <button type="submit">Register</button>
                <span>{message}</span>
            </form> 
            <span>Already have an account? <p2 onClick={() => navigate('/login')}> (Login) </p2></span>
        </div> 
    );
};

export default Register;
