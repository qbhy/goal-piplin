import {SubmitHandler, useForm} from "react-hook-form"
import {login} from "../services/auth.ts";
import {useNavigate} from "react-router-dom";

type Inputs = {
    username: string
    password: string
    remember: boolean
}

function Login() {
    const navigate = useNavigate()
    const {
        register,
        handleSubmit,
        formState: {errors},
    } = useForm<Inputs>()

    const onSubmit: SubmitHandler<Inputs> = (data) => login(data).then(res => {
        if (res.msg) {
            // todo
            alert(res.msg)
        } else {
            navigate('/dashboard')
        }
    })


    return (
        <div className='flex flex-col pt-[5rem] items-center bg-gray-100 min-h-screen'>
            <div className='text-2xl'>Goal-Piplin</div>
            <div className='mt-5 rounded p-3 bg-white'>
                <div className='text-center my-5'>登录 Goal-Piplin 账号</div>
                <form className='pt-3' onSubmit={handleSubmit(onSubmit)}>
                    <div className='flex flex-col'>
                        <input className='min-w-[20rem] border border-gray-200 rounded py-1 px-2 focus:outline-0'
                               placeholder='用户名或者邮箱' {...register("username", {required: true})} />
                        {errors.username && <span className='text-red-500 mt-1'>This field is required</span>}
                    </div>

                    <div className='mt-3 flex flex-col'>
                        <input className='min-w-[20rem] border border-gray-200 rounded py-1 px-2 focus:outline-0'
                               placeholder='密码' {...register("password", {required: true})} />
                        {errors.username && <span className='text-red-500 mt-1'>This field is required</span>}
                    </div>

                    <div className='flex justify-between mt-3'>
                        <label htmlFor="remember">
                            <input {...register('remember')} id='remember' name='remember'
                                   type="checkbox"/>
                            <span className='ml-2'>记住我</span>
                        </label>
                        <button type="submit" className='bg-blue-500 text-white py-1 px-3 rounded'>
                            登录
                        </button>
                    </div>
                </form>
            </div>
        </div>
    )
}

export default Login
