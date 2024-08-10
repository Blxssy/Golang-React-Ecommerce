import React, { useState } from 'react';
import api from '../services/api';
import { useNavigate } from 'react-router-dom';

const Register = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [message, setMessage] = useState('');
    const navigate = useNavigate();

    const handleRegister = async () => {
        try {
            const response = await api.post('/auth/register', { email, password });
            setMessage('Registration successful');
            navigate('/');
        } catch (error) {
            setMessage('Registration failed');
        }
    };

    return (
        <div>
            <h2>Register</h2>
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
            <button onClick={handleRegister}>Register</button>
            <p>{message}</p>
        </div>
    );
};

export default Register;
