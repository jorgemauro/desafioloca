// @flow
import classes from "./Navbar.module.scss";
import Link from "next/link";
const Navbar: React.FunctionComponent = () => {
  return (
    <nav
      className={`navbar navbar-expand-lg ${classes.root}`}
    >
      <div className={`container-fluid ${classes.navbarBody}`}>
        <Link href="/bank-accounts" as="/bank-accounts">
          <a className={`navbar-brand ${classes.navbarBrand}`} href="#">
            <img
              src="/img/icon_banco.png"
              alt=""
              className={classes.logoLocation}
            />
          </a>
        </Link>
          <div
            className={`collapse navbar-collapse ${classes.navbarRightRoot}`}
            id="navbarSupportedContent"
          >
            <ul className={`navbar-nav ml-auto ${classes.navbarRightBody}`}>
              <li className={`nav-item ${classes.bankAccountInfo}`}>
                <img
                  src="/img/icon_user.png"
                  alt=""
                  className={classes.iconUser}
                />
              </li>
            </ul>
          </div>
      </div>
    </nav>
  );
};

export default Navbar;
