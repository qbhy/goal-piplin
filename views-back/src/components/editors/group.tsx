import {FC} from "react";
import {Key} from "../../services/manage/keys.ts";
import {useForm} from "react-hook-form";
import classNames from "classnames";

export type KeyEditorProps = {
    defaultValue?: Key,
    onSubmit: (project: Key) => void,
    onClose: () => void
}


const GroupEditor: FC<KeyEditorProps> = ({defaultValue, onSubmit, onClose}) => {
    const {
        register,
        handleSubmit,
        formState: {errors/**/},
    } = useForm<Key>()


    return (
        <form onSubmit={handleSubmit(onSubmit)} className='block bg-white rounded-lg shadow'>
            <div className='py-3 text-base flex justify-between px-3 items-center'>
                <span>{defaultValue ? `编辑密钥 (${defaultValue.id})` : '新建密钥'}</span>
                <button className='p-2 hover:text-black hover:cursor-pointer' onClick={onClose}>x</button>
            </div>

            <div className='border-y px-5'>
                <div className='flex items-center py-3 w-[25rem]'>
                    <span>密钥名称</span>
                    <input className={classNames('ml-2 flex-1 focus:outline-0 rounded p-1', {
                        'border border-red-500': errors.name,
                    })} placeholder='输入密钥名称' {...register("name", {
                        required: true,
                        minLength: 2,
                        maxLength: 20
                    })} />
                </div>
            </div>

            <div className='p-3 flex justify-evenly gap-3'>
                <button className='inline-block text-center py-2 bg-blue-500 text-white px-5 rounded'
                        type='submit'>提交
                </button>
                <button onClick={onClose}
                        className='inline-block text-center py-2 border border-blue-500 text-blue-500 px-5 rounded'>取消
                </button>
            </div>
        </form>
    )
}

export default GroupEditor