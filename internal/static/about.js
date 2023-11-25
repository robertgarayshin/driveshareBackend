const slides = document.querySelectorAll('.slide');
let currentSlide = 0;

function showSlide(n) {
    slides[currentSlide].classList.remove('active');
    currentSlide = (n + slides.length) % slides.length;
    slides[currentSlide].classList.add('active');
}

document.getElementById('arrowLeft').addEventListener('click', () => {
    showSlide(currentSlide - 1);
});

document.getElementById('arrowRight').addEventListener('click', () => {
    showSlide(currentSlide + 1);
});


// Показ первого слайда при загрузке страницы
showSlide(0);