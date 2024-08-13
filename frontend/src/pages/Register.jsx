import React, { useState } from 'react';
import api from '../services/api';
import { useNavigate } from 'react-router-dom';
import './Register.css'; 

const Register = () => {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [message, setMessage] = useState('');
    const navigate = useNavigate();

    const handleRegister = async () => {
        try {
            await api.post('/auth/register', { username, email, password });
            setMessage('Registration successful');
            navigate('/');
        } catch (error) {
            setMessage('Registration failed');
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
                <p>{message}</p>
            </form> 
        </div> 
    );
};

export default Register;
