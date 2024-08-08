import React, { useState } from 'react'
import { Route, BrowserRouter as Router, Routes } from 'react-router-dom'
import './App.css'
import Auth from './pages/Auth'
import Home from './pages/Home'

function App() {
	const [auth, setAuth] = useState(null)

	const handleAuth = data => {
		setAuth(data)
	}

	return (
		<Router>
			<Routes>
				<Route path='/' element={<Home />} />
				<Route
					path='/register'
					element={<Auth onAuth={handleAuth} mode='register' />}
				/>
				<Route
					path='/login'
					element={<Auth onAuth={handleAuth} mode='login' />}
				/>
			</Routes>
		</Router>
	)
}

export default App
