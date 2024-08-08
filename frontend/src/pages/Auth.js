// src/pages/Auth.js
import React from 'react'
import { useLocation } from 'react-router-dom'
import AuthForm from '../components/AuthForm'

const Auth = ({ onAuth }) => {
	const location = useLocation()
	const isLogin = location.pathname === '/login'

	return <AuthForm isLogin={isLogin} onAuth={onAuth} />
}

export default Auth
