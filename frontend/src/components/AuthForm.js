// src/pages/Auth.js
import axios from 'axios'
import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import '../styles/AuthForm.css'

const Auth = ({ onAuth, mode }) => {
	const [email, setEmail] = useState('')
	const [password, setPassword] = useState('')
	const [error, setError] = useState('')
	const navigate = useNavigate()

	const handleSubmit = async e => {
		e.preventDefault()

		var url
		if (mode === 'register') {
			url = 'http://localhost:3001/api/auth/register'
		} else {
			url = 'http://localhost:3001/api/auth/login'
		}

		// mode === 'register'
		// 	? 'http://localhost:3001/api/auth/register'
		// 	: 'http://localhost:3001/api/auth/login'
		try {
			const response = await axios.post(url, { email, password })
			onAuth(response.data)
			if (mode === 'register') {
				navigate('/login') // Redirect to login page after successful registration
			} else {
				navigate('/') // Redirect to home page after successful login
			}
		} catch (error) {
			if (error.response && error.response.data) {
				setError('Invalid email or password')
			} else {
				setError('An unexpected error occurred')
			}
		}
	}

	return (
		<div className='auth-container'>
			<form onSubmit={handleSubmit} className='auth-form'>
				<h2>{mode === 'register' ? 'Register' : 'Login'}</h2>
				{error && <p className='error-message'>{error}</p>}
				<input
					type='email'
					placeholder='Email'
					value={email}
					onChange={e => setEmail(e.target.value)}
					className={error ? 'error-input' : ''}
					required
				/>
				<input
					type='password'
					placeholder='Password'
					value={password}
					onChange={e => setPassword(e.target.value)}
					className={error ? 'error-input' : ''}
					required
				/>
				<button type='submit'>
					{mode === 'register' ? 'Register' : 'Login'}
				</button>
			</form>
		</div>
	)
}

export default Auth
