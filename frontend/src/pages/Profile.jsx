import React from 'react';

const Profile = ({ user }) => {
    return (
        <div>
            <h2>Profile</h2>
            <p>Email: {user.email}</p>
            <p>Phone: {user.phone}</p>
            <img src={`${user.img}`} alt="Profile" />
        </div>
    );
};

export default Profile;
