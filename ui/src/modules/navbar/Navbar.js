

import React from 'react';
import { Navbar, Nav, Container, NavDropdown } from 'react-bootstrap';
import './Navbar.css'
import 'bootstrap/dist/css/bootstrap.min.css';

const MyNavbar = () => {
    const logo = require(`../../assets/cda-league-white.png`)
    return (
        <Navbar sticky="top" expand="lg" className="navbar navbar-color">
            <Navbar.Brand href="/">
                <img src={logo} alt="CDA League" className="navbar-logo" />
            </Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav navbar-toggler" className="navbar-hamburger" />
            <Navbar.Collapse id="esponsive-navbar-nav">
                <Nav className="me-auto nav-size">
                    <Nav.Link className="nav-link" href="rank">Tabla de Jugadores</Nav.Link>
                    <Nav.Link className="nav-link" href="historic">Historial</Nav.Link>
                    <Nav.Link className="nav-link" href="match">Armar Partido</Nav.Link>
                    <Nav.Link className="nav-link" href="paddle">Paddle</Nav.Link>
                </Nav>
            </Navbar.Collapse>
        </Navbar>
    );
};

export default MyNavbar;