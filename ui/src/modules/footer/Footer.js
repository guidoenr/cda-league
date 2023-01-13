import React from 'react';
import './Footer.css'

const Footer = () => {
    return (
        <footer>
            <p>Copyright © {new Date().getFullYear()} My App</p>
        </footer>
    );
}

export default Footer;
