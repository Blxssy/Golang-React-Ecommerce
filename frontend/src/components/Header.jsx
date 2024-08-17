import {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";
import api from "../services/api";
import '../styles/Header.css'
import Cookies from "js-cookie";
// import axios from 'axios';

const Header = () => {
    const [user, setUser] = useState()
    const navigate = useNavigate()

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await api.get('/auth/user-info');
                setUser(response.data)
            }catch (error) {
                console.log('Not authenticated')
            }
        }
        fetchUser();
    }, []);

    const handleLogout = () => {
        setUser(null);
        Cookies.remove("access_token")
        Cookies.remove("refresh_token")
    };

    return (
        <div className="header">
            {user ? (
                <div className="profile-section">
                    <h2>{user.username}</h2> 
                    <img
                        src={`${user.img}`}
                        alt="profile"
                        onClick={() => navigate('/profile')}
                        className="profile-avatar"
                    />
                    <button onClick={handleLogout}>Logout</button>
                </div>
            ) : (
                <div className="auth-buttons">
                    <button onClick={() => navigate('/register')}>Register</button>
                    <button onClick={() => navigate('/login')}>Login</button>
                </div>
            )}
        </div>
    );
};

export default Header;

