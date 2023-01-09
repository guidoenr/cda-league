import React, { useEffect } from "react";
import { useState } from "react";
import { NavLink, Link } from "react-router-dom";
import ThemeSwitcher from "./ThemeSwitcher";
import { HiMenuAlt3, HiOutlineUser } from "react-icons/hi";
import { MdOutlineClose } from "react-icons/md";
import { BsListTask } from "react-icons/bs";
import { GiTennisCourt } from "react-icons/gi";
import { FiLogOut } from "react-icons/fi";
import profileImage from "./assets/profile-example.jpg";

const Header = () => {
    const [nav, setNav] = useState(false);
    const [profileMenu, setProfileMenu] = useState(false);
    // @marcos, podes poner para que aca se vea el nombre del usuario? vamos a usar
    // el endpoint userGet para tener eso, no me gusta la foto del mono esa
    const [profileImg, setProfileImg] = useState(true);
    const [signOut, setSignOut] = useState(false);
    const [userName, setUsername] = useState("");

    const handleMenu = () => setNav(!nav);
    const handleProfileImg = () => setProfileImg(!profileImg);
    const handleProfileMenu = () => setProfileMenu(!profileMenu);
    const handleClose = () => setNav(!nav);

    useEffect(() => {
        (async () => {
            const response = await fetch("http://localhost:8080/auth/user", {
                headers: { "Content-Type": "application/json" },
                credentials: "include",
            });
            const content = await response.json();
            setUsername(content.name);
            setSignOut(true);
        })();
    });

    // Esto es lo que me querias explicar en tu casa:
    // Bienvenido {username}
    // Quedó piolaa!

    return (
        <header className="relative w-full h-[4.5rem] flex shadow-2xl z-50">
            <div className="container m-auto px-1 navbar flex justify-between">
                <div className="navbar-start w-auto">
                    <Link
                        to="/"
                        className="cursor-pointer font-secondary-font text-4xl text-primary"
                    >
                        Bienvenido {userName}
                    </Link>
                </div>

                <div className="flex justify-end">
                    <nav className="navbar-center hidden md:flex">
                        <ul className="menu menu-horizontal p-0 gap-8">
                            <li className="text-sm">
                                <NavLink
                                    to="/"
                                    className={({ isActive }) =>
                                        isActive
                                            ? "italic font-bold text-accent uppercase tracking-widest hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                                            : "italic font-medium uppercase tracking-widest hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                                    }
                                >
                                    Reservas
                                </NavLink>
                            </li>
                            <li className="text-sm">
                                <NavLink
                                    to="/misturnos"
                                    className={({ isActive }) =>
                                        isActive
                                            ? "italic font-bold text-accent uppercase tracking-widest hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                                            : "italic font-medium uppercase tracking-widest hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                                    }
                                >
                                    Mis Turnos
                                </NavLink>
                            </li>
                            <li className="text-sm">
                                <NavLink
                                    to="/profile"
                                    className={({ isActive }) =>
                                        isActive
                                            ? "italic font-bold text-accent uppercase tracking-widest hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                                            : "italic font-medium uppercase tracking-widest hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                                    }
                                >
                                    Mi Perfil
                                </NavLink>
                            </li>
                            <li className="flex items-center border-r border-l border-r-primary/30 border-l-primary/30 px-7">
                                <ThemeSwitcher className="" />
                            </li>
                            {signOut && (
                                <li className="">
                                    <Link
                                        to="/login"
                                        className="italic border-accent uppercase btn bg-primary/5 px-6 font-medium tracking-widest hover:border-accent/80 hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                                    >
                                        Entrar
                                    </Link>
                                </li>
                            )}
                        </ul>
                    </nav>

                    <div className="navbar-end relative">
                        <div className="flex items-center justify-end gap-4">
                            <div className="theme-container z-[90] md:hidden">
                                <ThemeSwitcher />
                            </div>

                            <div className="md:hidden mr-4 z-[90]" onClick={handleMenu}>
                                {!nav ? (
                                    <HiMenuAlt3 className="text-4xl cursor-pointer" />
                                ) : (
                                    <MdOutlineClose className="text-4xl cursor-pointer" />
                                )}
                            </div>
                        </div>
                        <div
                            className={
                                !nav
                                    ? "hidden"
                                    : "md:hidden absolute right-5 top-[4rem] h-[15rem] bg-base-100 w-[12rem] z-[60] rounded-box border border-primary/30"
                            }
                        >
                            <div className="flex flex-col w-full h-full justify-start relative">
                                <nav className="flex items-center w-full h-full">
                                    <ul className="menu menu-compact w-full bg-base-100 rounded-box">
                                        <li className="border-b border-b-primary/30 py-2 flex">
                                            <NavLink
                                                onClick={handleClose}
                                                to="/profile"
                                                className={({ isActive }) =>
                                                    isActive
                                                        ? "relative right-3 italic text-accent font-bold uppercase self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                        : "relative right-3 italic uppercase font-semibold self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                }
                                            >
                                                <HiOutlineUser className="text-2xl relative left-1" />
                                                Mi Perfil
                                            </NavLink>
                                        </li>
                                        <li className="border-b border-b-primary/30 py-2 flex">
                                            <NavLink
                                                onClick={handleClose}
                                                to="/"
                                                className={({ isActive }) =>
                                                    isActive
                                                        ? "relative right-2 italic text-accent font-bold uppercase self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                        : "relative right-2 italic uppercase font-semibold self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                }
                                            >
                                                <BsListTask className="text-2xl relative top-[.1rem] left-1" />
                                                Reservas
                                            </NavLink>
                                        </li>
                                        <li className="border-b border-b-primary/30 py-2 flex">
                                            <NavLink
                                                onClick={handleClose}
                                                to="/misturnos"
                                                className={({ isActive }) =>
                                                    isActive
                                                        ? "relative italic text-accent font-bold uppercase self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                        : "relative italic uppercase font-semibold self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                }
                                            >
                                                <GiTennisCourt className="text-2xl relative left-1" />
                                                Mis Turnos
                                            </NavLink>
                                        </li>
                                        <li className="py-2 flex">
                                            <NavLink
                                                onClick={() => {
                                                    handleClose();
                                                    setSignOut(true);
                                                    handleProfileImg(false);
                                                }}
                                                to="/"
                                                className={({ isActive }) =>
                                                    isActive
                                                        ? "relative right-7 italic text-red-700 font-bold uppercase self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                        : "relative right-7 italic uppercase text-red-700 font-semibold self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                }
                                            >
                                                <FiLogOut className="text-2xl relative left-2" />
                                                Salir
                                            </NavLink>
                                        </li>
                                    </ul>
                                </nav>
                            </div>
                        </div>
                    </div>

                    {profileImg && (
                        <div className="hidden md:flex md:ml-8 w-11 h-11 lg:w-12 lg:h-12 relative">
                            <div
                                className="z-[90] w-11 h-11 lg:w-12 lg:h-12 rounded-full cursor-pointer relative"
                                onClick={handleProfileMenu}
                            >
                                <img
                                    src={profileImage}
                                    alt="profile"
                                    className="w-11 h-11 rounded-full border border-accent lg:w-12 lg:h-12 relative bottom-[.1rem]"
                                />
                            </div>
                            <div
                                className={
                                    !profileMenu
                                        ? "hidden"
                                        : "hidden md:flex absolute right-0 top-[3.5rem] h-[16rem] w-[11rem] z-[60] rounded-box"
                                }
                            >
                                <div className="flex flex-col w-full h-full justify-start relative">
                                    <nav className="flex items-center w-full h-full">
                                        <ul className="menu menu-compact w-full bg-base-100 rounded-box border border-primary/30 shadow-lg">
                                            <li className="border-b border-b-primary/30 py-2 flex">
                                                <NavLink
                                                    onClick={handleProfileMenu}
                                                    to="/profile"
                                                    className={({ isActive }) =>
                                                        isActive
                                                            ? "relative right-3 italic text-accent font-bold uppercase self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                            : "relative right-3 italic uppercase font-semibold self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                    }
                                                >
                                                    <HiOutlineUser className="text-2xl relative left-1" />
                                                    Mi Perfil
                                                </NavLink>
                                            </li>
                                            <li className="border-b border-b-primary/30 py-2 flex">
                                                <NavLink
                                                    onClick={handleProfileMenu}
                                                    to="/"
                                                    className={({ isActive }) =>
                                                        isActive
                                                            ? "relative right-2 italic text-accent font-bold uppercase self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                            : "relative right-2 italic uppercase font-semibold self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                    }
                                                >
                                                    <BsListTask className="text-2xl relative top-[.1rem] left-1" />
                                                    Reservas
                                                </NavLink>
                                            </li>
                                            <li className="border-b border-b-primary/30 py-2 flex">
                                                <NavLink
                                                    onClick={handleProfileMenu}
                                                    to="/misturnos"
                                                    className={({ isActive }) =>
                                                        isActive
                                                            ? "relative italic text-accent font-bold uppercase self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                            : "relative italic uppercase font-semibold self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                    }
                                                >
                                                    <GiTennisCourt className="text-2xl relative left-1" />
                                                    Mis Turnos
                                                </NavLink>
                                            </li>
                                            <li className="py-2 flex">
                                                <NavLink
                                                    onClick={() => {
                                                        handleProfileMenu();
                                                        setSignOut(true);
                                                        handleProfileImg(false);
                                                    }}
                                                    to="/"
                                                    className={({ isActive }) =>
                                                        isActive
                                                            ? "relative right-7 italic text-red-700 font-bold uppercase self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                            : "relative right-7 italic uppercase text-red-700 font-semibold self-center text-sm tracking-widest hover:bg-transparent focus:bg-transparent"
                                                    }
                                                >
                                                    <FiLogOut className="text-2xl relative left-2" />
                                                    Salir
                                                </NavLink>
                                            </li>
                                        </ul>
                                    </nav>
                                </div>
                            </div>
                        </div>
                    )}
                </div>
            </div>
        </header>
    );
};

export default Header;