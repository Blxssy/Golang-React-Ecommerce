import React, { useEffect, useState, useRef } from 'react';
import '../styles/slider.css';

const Slider = () => {
  const [products, setProducts] = useState([]);
  const [step, setStep] = useState(0);
  const [offset, setOffset] = useState(0);
  const itemsRef = useRef(null);

  useEffect(() => {
    fetch('http://localhost:3001/api/products')
      .then(response => response.json())
      .then(data => {
        // Клонируем первый и последний продукт для бесконечного слайдера
        setProducts([data[data.length - 1], ...data, data[0]]);
      })
      .catch(error => console.error('Error loading products:', error));
  }, []);

  useEffect(() => {
    if (products.length > 0) {
      burgerSlider();
    }
  }, [products]);

  const burgerSlider = () => {
    const lastProduct = products[products.length - 1];
    const currentProduct = products[step];
    const nextProduct = products[step + 1];

    itemsRef.current.innerHTML = ''; // Очищаем контейнер

    const createDiv = (product, leftOffset) => {
      const div = document.createElement('div');
      div.className = 'item';
      div.style.left = `${leftOffset * 100}px`;
      div.textContent = product.name; // Или другие данные продукта
      return div;
    };

    itemsRef.current.appendChild(createDiv(lastProduct, -1));
    itemsRef.current.appendChild(createDiv(currentProduct, 0));
    itemsRef.current.appendChild(createDiv(nextProduct, 1));

    setOffset(1);
  };

  const burgerSliderL = () => {
    let newStep = step;

    if (newStep === products.length - 1) {
      newStep = 1;
    } else if (newStep === products.length - 2) {
      newStep = 0;
    } else {
      newStep += 2;
    }

    const currentProduct = products[newStep];
    const newDiv = createDiv(currentProduct, offset * 100);

    itemsRef.current.appendChild(newDiv);

    if (newStep === 0) {
      newStep = products.length - 1;
    } else {
      newStep -= 1;
    }

    setStep(newStep);
    setOffset(1);
  };

  const left = () => {
    const sliderItems = document.querySelectorAll('.item');
    let offset2 = -1;

    sliderItems.forEach(item => {
      item.style.left = `${offset2 * 100 - 100}px`;
      offset2++;
    });

    setTimeout(() => {
      sliderItems[0].remove();
      burgerSliderL();
    }, 600);
  };

  const burgerSliderR = () => {
    let newStep = step;

    if (newStep === 0) {
      newStep = products.length - 2;
    } else if (newStep === 1) {
      newStep = products.length - 1;
    } else {
      newStep -= 2;
    }

    const currentProduct = products[newStep];
    const newDiv = createDiv(currentProduct, -1);

    itemsRef.current.insertBefore(newDiv, itemsRef.current.firstElementChild);

    if (newStep === products.length - 1) {
      newStep = 0;
    } else {
      newStep += 1;
    }

    setStep(newStep);
  };

  const right = () => {
    const sliderItems = document.querySelectorAll('.item');
    let offset2 = sliderItems.length - 1;

    for (let i = sliderItems.length - 1; i >= 0; i--) {
      sliderItems[i].style.left = `${offset2 * 100}px`;
      offset2--;
    }

    setTimeout(() => {
      sliderItems[sliderItems.length - 1].remove();
      burgerSliderR();
    }, 600);
  };

  const createDiv = (product, leftOffset) => {
    const div = document.createElement('div');
    div.className = 'item';
    div.style.left = `${leftOffset * 100}px`;
    div.textContent = product.name; // или другое свойство продукта
    return div;
  };

  return (
    <div className="slider-container">
      <a className="arrow" id="leftArrow" href="#" onClick={left}>←</a>
      <div className="items" id="items" ref={itemsRef}></div>
      <a className="arrow" id="rightArrow" href="#" onClick={right}>→</a>
    </div>
  );
};

export default Slider;
