import { useEffect, useState } from 'react'
import '../styles/ProductsList.css'
import { ProductCard } from './ProductCard'

export const ProductsList = () => {
	const [products, setProducts] = useState([])

	useEffect(() => {
		fetch('http://localhost:3001/api/products')
			.then(response => response.json())
			.then(data => setProducts(data))
			.catch(error => console.error('Error fetching users:', error))
	}, [])

	return (
		<div className='products-container'>
			<div className='products-list'>
				{products.map(product => (
					<ProductCard key={product.ID} product={product} />
				))}
			</div>
		</div>
	)
}
