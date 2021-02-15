import classes from './Footer.module.scss';
console.log(classes);
const Footer: React.FunctionComponent = (props) => {
  return (
    <footer className={classes.root}>
      © desafioloca - Todos os direitos reservados.
    </footer>
  );
};

export default Footer;