import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import Profile from './pages/Profile';
import Header from "./components/Header";
import BottomMenu from "./components/BottomMenu";
import './App.css';
import Register from "./pages/Register";
import Login from "./pages/Login";
import ProductPage from './pages/ProductPage';
import ProductList from './pages/ProductList';
import HomeTest from './pages/Home_test';
import Cart from './pages/Cart';

function App() {
	return (

			<Router>
				<div className="App">
					<Header />
					<BottomMenu />
					<Routes>
						<Route path="/" element={<Home />} />
						<Route path="/test" element={<HomeTest />} />
						<Route path="/register" element={<Register/>}/>
						<Route path="/login" element={<Login/>}/>
						<Route path="/profile" element={<Profile />} />
						<Route path="/products" element={<ProductList />} />
						<Route path="/products/:id" element={<ProductPage />} /> 
						<Route path="/cart" element={<Cart />} />
					</Routes>
				</div>
			</Router>

	);
}

export default App;
