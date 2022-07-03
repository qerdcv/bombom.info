import styles from 'components/Loader/loader.module.scss';

export const Loader = () => {
    return (
        <div className={styles.loader}>
            <div></div>
            <div></div>
            <div></div>
            <div></div>
        </div>
    );
};
