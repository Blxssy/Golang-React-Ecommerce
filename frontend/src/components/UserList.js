// src/components/UserList.js
import React, { useEffect, useState } from 'react'
import UserCard from './UserCard'
import './UserList.css'

const UserList = () => {
	const [users, setUsers] = useState([])

	useEffect(() => {
		fetch('http://localhost:3001/api/users')
			.then(response => response.json())
			.then(data => setUsers(data))
			.catch(error => console.error('Error fetching users:', error))
	}, [])

	return (
		<div className='user-container'>
			<div className='user-list'>
				{users.map(user => (
					<UserCard key={user.ID} user={user} />
				))}
			</div>
		</div>
	)
}

export default UserList
