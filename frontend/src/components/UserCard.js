// src/components/UserCard.js
import React from 'react'
import '../styles/UserCard.css'

const UserCard = ({ user }) => {
	return (
		<div className='user-card'>
			<img
				src={user.AvatarPath}
				alt={`${user.username}'s avatar`}
				className='avatar'
			/>
			<h2>{user.username}</h2>
			<p>Email: {user.email}</p>
			<p>Phone: {user.phone}</p>
		</div>
	)
}

export default UserCard
