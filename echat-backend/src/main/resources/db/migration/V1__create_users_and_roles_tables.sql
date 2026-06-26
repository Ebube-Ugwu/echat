create table if not exists roles (
    id          bigserial primary key,
    name        varchar(50)  not null,
    description varchar(255),
    built_in    boolean      not null default false,
    constraint uk_role_name unique (name)
);

create table if not exists role_permissions (
    role_id    bigint      not null,
    permission varchar(60) not null,
    constraint fk_role_permissions_role
        foreign key (role_id) references roles(id) on delete CASCADE,
    constraint uk_role_permission unique (role_id, permission)
);

create table if not exists users (
  id bigserial primary key,
  email text not null,
  username text not null,
  password text not null,
  password_changed_at timestamptz not null default now(),
  status text not null default 'ACTIVE',
  auth_provider text not null default 'EMAIL',
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),
  constraint uq_user_email unique (email),
  constraint user_status_check check (status in ('ACTIVE', 'INACTIVE', 'SUSPENDED', 'PENDING_VERIFY')),
  constraint user_auth_provider_check check (auth_provider in ('EMAIL', 'APPLE', 'GOOGLE'))
);

create table if not exists user_roles (
    user_id bigint not null,
    role_id bigint not null,
    primary key (user_id, role_id),
    constraint fk_user_roles_user
        foreign key (user_id) references users(id) on delete cascade,
    constraint fk_user_roles_role
        foreign key (role_id) references roles(id) on delete cascade
);
