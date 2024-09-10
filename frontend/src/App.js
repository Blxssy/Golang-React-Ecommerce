// import React from 'react';
// import { BrowserRouter as Router, Route, Routes, useLocation } from 'react-router-dom';
// import Home from './pages/Home';
// import Profile from './pages/Profile';
// import Header from "./components/Header";
// import BottomMenu from "./components/BottomMenu";
// import './App.css';
// import Register from "./pages/Register";
// import Login from "./pages/Login";
// import ProductPage from './pages/ProductPage';
// import ProductList from './pages/ProductList';
// import HomeTest from './pages/Home_test';
// import Cart from './pages/Cart';
// import Authorize from './pages/Authorize';

// function App() {
// 	const location = useLocation();
// 	const shouldShowHeaderAndMenu = !['/authorize', '/register', '/login'].includes(location.pathname);

// 	return (

// 			<Router>
// 				<div className="App">
// 					{shouldShowHeaderAndMenu && <Header />}
// 					{shouldShowHeaderAndMenu && <BottomMenu />}

// 					<Header />
// 					<BottomMenu />
// 					<Routes>
// 						<Route path="/" element={<Home />} />
// 						<Route path="/test" element={<HomeTest />} />
// 						<Route path="/authorize" element={<Authorize/>}/>
// 						<Route path="/register" element={<Register/>}/>
// 						<Route path="/login" element={<Login/>}/>
// 						<Route path="/profile" element={<Profile />} />
// 						<Route path="/products" element={<ProductList />} />
// 						<Route path="/products/:id" element={<ProductPage />} /> 
// 						<Route path="/cart" element={<Cart />} />
// 					</Routes>
// 				</div>
// 			</Router>

// 	);
// }

// export default App;


import { BrowserRouter as Router, Routes, Route, useLocation } from 'react-router-dom';
import Header from './components/Header';
import BottomMenu from './components/BottomMenu';
import Home from './pages/Home';
// import HomeTest from './pages/HomeTest';
import Authorize from './pages/Authorize';
import Register from './pages/Register';
import Login from './pages/Login';
import Profile from './pages/Profile';
import ProductList from './pages/ProductList';
import ProductPage from './pages/ProductPage';
import Cart from './pages/Cart';

// Компонент для условного отображения Header и BottomMenu
function AppContent() {
  const location = useLocation(); // Получаем текущий путь

  // Условие для отображения Header и BottomMenu
  const shouldShowHeaderAndMenu = !['/authorize', '/register', '/login'].includes(location.pathname);

  return (
    <div className="App">
      {/* Отображаем Header и BottomMenu, если путь не /authorize, /register или /login */}
      {shouldShowHeaderAndMenu && <Header />}
      {shouldShowHeaderAndMenu && <BottomMenu />}

      <Routes>
        <Route path="/" element={<Home />} />
        {/* <Route path="/test" element={<HomeTest />} /> */}
        <Route path="/authorize" element={<Authorize />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
        <Route path="/profile" element={<Profile />} />
        <Route path="/products" element={<ProductList />} />
        <Route path="/products/:id" element={<ProductPage />} />
        <Route path="/cart" element={<Cart />} />
      </Routes>
    </div>
  );
}

function App() {
  return (
    <Router>
      <AppContent />
    </Router>
  );
}

export default App;
