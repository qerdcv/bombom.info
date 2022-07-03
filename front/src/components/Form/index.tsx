import { useForm } from 'react-hook-form';

import styles from 'components/Form/form.module.scss';
import { useHttp } from 'hooks/http';

interface IJoinRequest {
    userTag:      string;
    telegramName: string;
}

const Required = () => {
    return <span style={{color: 'red'}}>*</span>
}

export const JoinRequestForm = () => {
    const { register, handleSubmit, formState: { errors }} = useForm<IJoinRequest>();
    const { request, isLoading, error } = useHttp()

    const onSubmit = async (data: IJoinRequest) => {
        request('post', '/api/v1/join', data)
    }

    console.log(isLoading, error)

    return (
        <form onSubmit={handleSubmit(onSubmit)} className={styles.form}>
            <div className={styles.formControl}>
                <label htmlFor="userTag">User tag<Required />:</label>
                <input id='user_tag' autoComplete='off' className={styles.input} {...register('userTag', {
                    required: {
                        value: true,
                        message: 'User tag is required.'
                    },
                    pattern: {
                        value: /#([a-zA-Z]|\d)*/,
                        message: 'User tag must start with # and consist only with letters and numbers.'
                    }
                })}/>
                {errors.userTag && (
                    <small className={styles.error}>{errors.userTag.message}</small>
                )}
            </div>
            <div className={styles.formControl}>
                <label htmlFor="telegramName">Telegram<Required />:</label>
                <input id='telegramName' autoComplete='off' className={styles.input} {...register('telegramName', {
                    required: {
                        value: true,
                        message: 'Telegram username is required.'
                    },
                    pattern: {
                        value: /@([a-zA-Z]|\d)*/,
                        message: 'Telegram username must start with @ and consist only with letters and numbers.'
                    },
                })}/>
                {errors.telegramName && (
                    <small className={styles.error}>{errors.telegramName.message}</small>
                )}
            </div>
            <div className={styles.formControl}>
                <input type='submit' value='Request to join' className={styles.btn} disabled={isLoading} />
            </div>
            {error && <small className={styles.error}> {error.message} </small>}
        </form>
    );
}
