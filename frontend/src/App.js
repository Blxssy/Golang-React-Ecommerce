import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import Profile from './pages/Profile';
import Header from "./components/Header";
import './App.css';
import Register from "./pages/Register";
import Login from "./pages/Login";
import ProductPage from './pages/ProductPage';
import ProductList from './pages/ProductList';
import Home_test from './pages/Home_test';

function App() {
	return (

			<Router>
				<div className="App">
					<Header />
					<Routes>
						<Route path="/" element={<Home />} />
						<Route path="/test" element={<Home_test />} />
						<Route path="/register" element={<Register/>}/>
						<Route path="/login" element={<Login/>}/>
						<Route path="/profile" element={<Profile />} />
						<Route path="/products" element={<ProductList />} />
						<Route path="/products/:id" element={<ProductPage />} /> 
					</Routes>
				</div>
			</Router>

	);
}

export default App;
