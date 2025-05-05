import { useOutside } from "hooks/useOutside";
import { FC, useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import Meta from "utils/Meta";
import { IAuthFields } from "./login-form.interface";
import { useAuth } from "hooks/useAuth";
import styles from "./LoginForm.module.scss";
import { FaUserCircle } from "react-icons/fa";
import { validEmail } from "./login-auth.constants";
import Button from "../../Button/Button";
import Field from "../../Field/Field";
import UserAvatar from "../../user-avatar/UserAvatar";

const LoginForm: FC = () => {
  const { ref, setIsShow, isShow } = useOutside(false);
  const [type, setType] = useState<"login" | "register">("login");
  const {
    register,
    formState: { errors },
    handleSubmit,
  } = useForm<IAuthFields>({
    mode: "onChange",
  });

  const { user, setUser } = useAuth();
  const onSubmit: SubmitHandler<IAuthFields> = (data) => {
    if (type == "login")
      setUser({
        id: 1,
        email: "test@test.ru",
        avatarPath: "/clapper.ico",
        name: "Dev",
      });
    // else if (type=='register') registerSync(data)
  };
  return (
    <div className={styles.wrapper} ref={ref}>
      {user ? (
        <UserAvatar />
      ) : (
        <button
          style={{
            background: "none",
          }}
          className={styles.button}
          onClick={() => setIsShow(!isShow)}
        >
          <FaUserCircle fill="#A4A4A4" className={styles.userIcon} />
        </button>
      )}

      {isShow && (
        <form onSubmit={handleSubmit(onSubmit)} className={styles.form}>
          <Field
            type="email"
            placeholder="Email"
            error={errors.email}
            {...register("email", {
              required: "Email is required",
              pattern: {
                value: validEmail,
                message: "Please enter a valid email address",
              },
            })}
          />
          <Field
            type="password"
            placeholder="password"
            error={errors.password}
            {...register("password", {
              required: "password is required",
              minLength: {
                value: 6,
                message: "Min length should more 6 symbols",
              },
            })}
          />
          <div className={"mt-5 mb-1 text-center"}>
            <Button onClick={() => setType("login")}>Login</Button>
          </div>
          <button onClick={() => setType("register")}>Register</button>
        </form>
      )}
    </div>
    //   <div className="text-black">LoginForm</div>
  );
};
export default LoginForm;
