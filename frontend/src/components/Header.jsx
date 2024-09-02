import {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";
import api from "../services/api";
import '../styles/Header.css'
import Cookies from "js-cookie";
import { IoIosHome } from "react-icons/io";
import { IoIosLogOut } from "react-icons/io";

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
                        <div className="auth-buttons">
                    <button-logout onClick={handleLogout}>
                        Logout 
                        <IoIosLogOut />
                    </button-logout>
                    <button-home onClick={() => navigate('/ ')}> Home <IoIosHome /></button-home>
                        </div>
                        <h2>{user.username}</h2> 
                    <img
                        // src="https://bigpicture.ru/wp-content/uploads/2021/10/bigpicture_ru_duejn-skala-dzhonson-zhestko-mem.jpg"
                        src={`${user.img}`}
                        alt="profile"
                        onClick={() => navigate('/profile')}
                        className="profile-avatar"
                    />
                </div>
            ) : (
                <div className="auth-buttons">
                    <button onClick={() => navigate('/register')}>Register</button>
                    <button onClick={() => navigate('/login')}>Login</button>
                    <button-home onClick={() => navigate('/ ')}> Home <IoIosHome /></button-home>
                </div>
            )}
        </div>
    );
};

export default Header;

