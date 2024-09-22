import {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";
import api from "../services/api";
import '../styles/Header.css'
import Cookies from "js-cookie";
// import { IoIosHome } from "react-icons/io";
// import { IoIosLogOut } from "react-icons/io";
import { IoIosHome, IoIosLogOut } from 'react-icons/io';

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
              <div className="header-logged">
                <div className="header-container">
                  <h2 className="header-title">{user.username}</h2>
                  <img
                    src={`${user.img}`}
                    alt="profile"
                    onClick={() => navigate('/profile')}
                    className="profile-avatar"
                  />
                </div>
              </div>
              <div className="auth-buttons">
                <button className="button-logout" onClick={handleLogout}>
                  Logout&nbsp;<IoIosLogOut style={{ position: 'relative', top: '-2px' }} />
                </button>
                <button className="button-home" onClick={() => navigate('/')}>
                  Home&nbsp;<IoIosHome style={{ position: 'relative', top: '-2px' }} />
                </button>
              </div>
            </div>
          ) : (
            <div className="auth-buttons">
              <button onClick={() => navigate('/authorize')}>Register</button>
              <button onClick={() => navigate('/authorize')}>Login</button>
              <button className="button-home-nl" onClick={() => navigate('/')}>
                Home&nbsp;<IoIosHome style={{ position: 'relative', top: '-2px' }} />
              </button>
            </div>
          )}
        </div>
      );
};

export default Header;

