import React from 'react'
import '../styles/ProductCard.css'

export const ProductCard = ({ product }) => {
	return (
		<div className='product-card'>
			<img src={product.image} alt={`${product.name}`} className='img' />
			<div className='price'>
				<p>{product.price}$</p>
			</div>
			<h2>{product.name}</h2>
			{/* <p>{product.description}</p> */}
		</div>
	)
}
