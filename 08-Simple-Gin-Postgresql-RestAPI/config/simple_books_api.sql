PGDMP         #                {            simple_books_api    15.2    15.2     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16410    simple_books_api    DATABASE     �   CREATE DATABASE simple_books_api WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
     DROP DATABASE simple_books_api;
                postgres    false            �            1259    16412    books    TABLE     �   CREATE TABLE public.books (
    id integer NOT NULL,
    title character varying(100) NOT NULL,
    author character varying(50) NOT NULL,
    description text NOT NULL
);
    DROP TABLE public.books;
       public         heap    postgres    false            �            1259    16411    books_id_seq    SEQUENCE     �   CREATE SEQUENCE public.books_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.books_id_seq;
       public          postgres    false    215            �           0    0    books_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;
          public          postgres    false    214            e           2604    16415    books id    DEFAULT     d   ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);
 7   ALTER TABLE public.books ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    214    215    215            �          0    16412    books 
   TABLE DATA           ?   COPY public.books (id, title, author, description) FROM stdin;
    public          postgres    false    215   �
       �           0    0    books_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.books_id_seq', 6, true);
          public          postgres    false    214            g           2606    16419    books books_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.books DROP CONSTRAINT books_pkey;
       public            postgres    false    215            �   �   x�M�A
�0EדS�JA�]�"�Fݹ�6�$���i�^ߢ���6p���['yV���E'f�^^ǌ'���hHΤ��H3k��͇W�S-�P`RqJ��+������uDi?j����l�,e�p��){�߅���<Jc�7>�     