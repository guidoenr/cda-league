import React from 'react';
import { Navbar, Nav, Container, NavDropdown } from 'react-bootstrap';
import './Navbar.css'
import 'bootstrap/dist/css/bootstrap.min.css';

const MyNavbar = () => {
    const logo = require(`../assets/carmen-league.png`)
    return (
        <Navbar className="mynavbar fixed-top" expand="lg" >
                <div className="navbar-logo">
                    <Navbar.Brand href="/">
                        <img src={logo} alt="CDA League" width="180" height="120" className="d-inline-block align-top nav" />
                    </Navbar.Brand>
                </div>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav hamburger">
                    <Nav className="me-auto nav-size">
                        <Nav.Link href="/">Home</Nav.Link>
                        <Nav.Link href="players">Jugadores</Nav.Link>
                        <Nav.Link href="rank">Tabla/Rank</Nav.Link>
                        <Nav.Link href="match">Partido</Nav.Link>
                        <Nav.Link href="paddle">Paddle</Nav.Link>
                    </Nav>
                </Navbar.Collapse>
        </Navbar>
    );
};

export default MyNavbar;