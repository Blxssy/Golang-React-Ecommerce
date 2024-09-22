import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Cookies from 'js-cookie';
import api from '../services/api';
import '../styles/Authorize.css';

const Authorize = () => {
  const [isRegister, setIsRegister] = useState(false);
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const navigate = useNavigate();

  // Регистрация
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
        setMessage('Registration failed, try another email or password');
      }
    }
  };

  // Вход
  const handleLogin = async () => {
    try {
      const response = await api.post('/auth/login', { email, password });
      const { access_token, refresh_token } = response.data;

      Cookies.set('access_token', access_token);
      Cookies.set('refresh_token', refresh_token);

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

  const handleSubmit = (e) => {
    e.preventDefault();
    if (isRegister) {
      handleRegister();
    } else {
      handleLogin();
    }
  };

  return (
    <div className={`container ${isRegister ? 'active' : ''}`} id="container">
      {/* Форма регистрации */}
      <div className="form-container sign-up">
        <form onSubmit={handleSubmit}>
          <h1 style = {{ color: "363636" }}>Create Account</h1>
          <input
            type="text"
            placeholder="Name"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <input
            type="email"
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
          <button type="submit">Sign Up</button>
          <span>{message}</span>
        </form>
      </div>

      {/* Форма входа */}
      <div className="form-container sign-in">
        <form onSubmit={handleSubmit}>
          <h1 style = {{ color: "363636" }}>Sign In</h1>
          <input
            type="email"
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
          <a href="#">Forget Your Password?</a>
          <button type="submit">Login</button>
          <span>{message}</span>
        </form>
      </div>

      {/* Переключатель между регистрацией и входом */}
      <div className="toggle-container">
        <div className="toggle">
          <div className="toggle-panel toggle-left">
            <h1>Welcome Back!</h1>
            <p>Enter your personal details to use all of the site features</p>
            <button className="hidden" id="login" onClick={() => setIsRegister(false)}>Sign In</button>
          </div>
          <div className="toggle-panel toggle-right">
            <h1>Hello, Friend!</h1>
            <p>Register with your personal details to use all of the site features</p>
            <button className="hidden" id="register" onClick={() => setIsRegister(true)}>Sign Up</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Authorize;
