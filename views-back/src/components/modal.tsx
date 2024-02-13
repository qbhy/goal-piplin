import {FC, PropsWithChildren} from "react";
import classNames from "classnames";

type ModalProps = {
    visible: boolean
}

const Modal: FC<ModalProps & PropsWithChildren> = ({visible, children}) => {
    return (
        <div
            className={classNames('fixed top-0 left-0 right-0 bottom-0 bg-gray-800/75 z-50 flex justify-center items-center', {
                'hidden': !visible,
            })}>
            {children}
        </div>
    )
}

export default Modal