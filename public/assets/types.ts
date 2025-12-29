// types.ts

export type Identifier = string;

export type Platform = 'web' | 'mobile' | 'desktop';

export type Asset = string;

export type Url = string;

export type Callback = () => void;

export type CallbackWithArgs<T> = (arg: T) => void;

export type Resolved<T> = T | Promise<T>;

export type Error<T> = { error: T };

export type Optional<T> = T | null;

export type ArrayOf<T> = T[];

export type ObjectOf<T> = { [key: string]: T };

export type MappedObject<T, U> = { [key in keyof T]: U };

export type IndexedObject<T> = { [key: string]: T };

export type Enum<T> = T[keyof T];

export type MappedEnum<T, U> = T[keyof T] extends Enum<Enum<U>> ? U : never;

export type EnumKeys<T> = T[keyof T];

export type Never = never;

export type Undefined = undefined;