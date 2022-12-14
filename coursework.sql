PGDMP          %            
    z            MiDB    14.5    14.5 N    a           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            b           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            c           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            d           1262    18359    MiDB    DATABASE     c   CREATE DATABASE "MiDB" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Russian_Russia.1251';
    DROP DATABASE "MiDB";
                postgres    false            ?            1259    18385    Category    TABLE     J   CREATE TABLE public."Category" (
    "Name" character varying NOT NULL
);
    DROP TABLE public."Category";
       public         heap    postgres    false            ?            1259    18399    Comments    TABLE     ?   CREATE TABLE public."Comments" (
    "ID" bigint NOT NULL,
    "Product" bigint NOT NULL,
    "User" character varying NOT NULL,
    "Messenger" text NOT NULL
);
    DROP TABLE public."Comments";
       public         heap    postgres    false            ?            1259    18658    DelUsers    TABLE     k   CREATE TABLE public."DelUsers" (
    "Email" character varying NOT NULL,
    "Reason" character varying
);
    DROP TABLE public."DelUsers";
       public         heap    postgres    false            ?            1259    18430    MainDiscount    TABLE     7  CREATE TABLE public."MainDiscount" (
    "ID" integer NOT NULL,
    "Product" character varying NOT NULL,
    "Name" bigint NOT NULL,
    "Characteristic" character varying NOT NULL,
    "NewPrice" character varying NOT NULL,
    "OldPrice" character varying NOT NULL,
    "Image" character varying NOT NULL
);
 "   DROP TABLE public."MainDiscount";
       public         heap    postgres    false            ?            1259    18429    MainDiscount_ID_seq    SEQUENCE     ?   CREATE SEQUENCE public."MainDiscount_ID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public."MainDiscount_ID_seq";
       public          postgres    false    222            e           0    0    MainDiscount_ID_seq    SEQUENCE OWNED BY     Q   ALTER SEQUENCE public."MainDiscount_ID_seq" OWNED BY public."MainDiscount"."ID";
          public          postgres    false    221            ?            1259    18406    Manager    TABLE     ?   CREATE TABLE public."Manager" (
    "Email" character varying NOT NULL,
    "Name" character varying NOT NULL,
    "Role" character varying NOT NULL,
    "Password" character varying
);
    DROP TABLE public."Manager";
       public         heap    postgres    false            ?            1259    18665 	   Periphery    TABLE        CREATE TABLE public."Periphery" (
    "Article" bigint NOT NULL,
    "Series" character varying NOT NULL,
    "Name" character varying NOT NULL,
    "Price" double precision NOT NULL,
    "Count" bigint NOT NULL,
    "Image" character varying NOT NULL,
    "Discription" text NOT NULL
);
    DROP TABLE public."Periphery";
       public         heap    postgres    false            ?            1259    18643    Posts    TABLE     ?   CREATE TABLE public."Posts" (
    "Id" integer NOT NULL,
    "Image" character varying NOT NULL,
    "Title" character varying NOT NULL,
    "Text" character varying NOT NULL,
    "Data" character varying NOT NULL
);
    DROP TABLE public."Posts";
       public         heap    postgres    false            ?            1259    18642    Posts_Id_seq    SEQUENCE     ?   CREATE SEQUENCE public."Posts_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public."Posts_Id_seq";
       public          postgres    false    225            f           0    0    Posts_Id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public."Posts_Id_seq" OWNED BY public."Posts"."Id";
          public          postgres    false    224            ?            1259    18421    PreOrder    TABLE       CREATE TABLE public."PreOrder" (
    "ID" integer NOT NULL,
    "Date" character varying NOT NULL,
    "Product" character varying NOT NULL,
    "Benefit" character varying NOT NULL,
    "Gift" character varying NOT NULL,
    "Image" character varying NOT NULL
);
    DROP TABLE public."PreOrder";
       public         heap    postgres    false            ?            1259    18420    PreOrder_ID_seq    SEQUENCE     ?   CREATE SEQUENCE public."PreOrder_ID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public."PreOrder_ID_seq";
       public          postgres    false    220            g           0    0    PreOrder_ID_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public."PreOrder_ID_seq" OWNED BY public."PreOrder"."ID";
          public          postgres    false    219            ?            1259    18392    ProductCategory    TABLE     t   CREATE TABLE public."ProductCategory" (
    "Product" bigint NOT NULL,
    "Category" character varying NOT NULL
);
 %   DROP TABLE public."ProductCategory";
       public         heap    postgres    false            ?            1259    18361    Products    TABLE     !  CREATE TABLE public."Products" (
    "Article" integer NOT NULL,
    "Series" character varying NOT NULL,
    "Name" character varying NOT NULL,
    "Price" double precision NOT NULL,
    "Count" integer NOT NULL,
    "Image" character varying NOT NULL,
    "Description" text NOT NULL
);
    DROP TABLE public."Products";
       public         heap    postgres    false            ?            1259    18360    Products_Article_seq    SEQUENCE     ?   CREATE SEQUENCE public."Products_Article_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public."Products_Article_seq";
       public          postgres    false    210            h           0    0    Products_Article_seq    SEQUENCE OWNED BY     S   ALTER SEQUENCE public."Products_Article_seq" OWNED BY public."Products"."Article";
          public          postgres    false    209            ?            1259    18377    Purchase    TABLE       CREATE TABLE public."Purchase" (
    "ID" integer NOT NULL,
    "User" character varying,
    "Product" bigint NOT NULL,
    "Count" integer NOT NULL,
    "Price" double precision NOT NULL,
    "Date" character varying NOT NULL,
    "TotalPrice" double precision NOT NULL
);
    DROP TABLE public."Purchase";
       public         heap    postgres    false            ?            1259    18376    Purchase_ID_seq    SEQUENCE     ?   CREATE SEQUENCE public."Purchase_ID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public."Purchase_ID_seq";
       public          postgres    false    213            i           0    0    Purchase_ID_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public."Purchase_ID_seq" OWNED BY public."Purchase"."ID";
          public          postgres    false    212            ?            1259    18413    Role    TABLE     d   CREATE TABLE public."Role" (
    "Name" character varying NOT NULL,
    "Access" bigint NOT NULL
);
    DROP TABLE public."Role";
       public         heap    postgres    false            ?            1259    18438    Series    TABLE     H   CREATE TABLE public."Series" (
    "Name" character varying NOT NULL
);
    DROP TABLE public."Series";
       public         heap    postgres    false            ?            1259    18369    Users    TABLE     ?   CREATE TABLE public."Users" (
    "Email" character varying NOT NULL,
    "Password" character varying NOT NULL,
    "Name" character varying NOT NULL
);
    DROP TABLE public."Users";
       public         heap    postgres    false            ?           2604    18433    MainDiscount ID    DEFAULT     x   ALTER TABLE ONLY public."MainDiscount" ALTER COLUMN "ID" SET DEFAULT nextval('public."MainDiscount_ID_seq"'::regclass);
 B   ALTER TABLE public."MainDiscount" ALTER COLUMN "ID" DROP DEFAULT;
       public          postgres    false    222    221    222            ?           2604    18646    Posts Id    DEFAULT     j   ALTER TABLE ONLY public."Posts" ALTER COLUMN "Id" SET DEFAULT nextval('public."Posts_Id_seq"'::regclass);
 ;   ALTER TABLE public."Posts" ALTER COLUMN "Id" DROP DEFAULT;
       public          postgres    false    224    225    225            ?           2604    18424    PreOrder ID    DEFAULT     p   ALTER TABLE ONLY public."PreOrder" ALTER COLUMN "ID" SET DEFAULT nextval('public."PreOrder_ID_seq"'::regclass);
 >   ALTER TABLE public."PreOrder" ALTER COLUMN "ID" DROP DEFAULT;
       public          postgres    false    220    219    220            ?           2604    18364    Products Article    DEFAULT     z   ALTER TABLE ONLY public."Products" ALTER COLUMN "Article" SET DEFAULT nextval('public."Products_Article_seq"'::regclass);
 C   ALTER TABLE public."Products" ALTER COLUMN "Article" DROP DEFAULT;
       public          postgres    false    209    210    210            ?           2604    18380    Purchase ID    DEFAULT     p   ALTER TABLE ONLY public."Purchase" ALTER COLUMN "ID" SET DEFAULT nextval('public."Purchase_ID_seq"'::regclass);
 >   ALTER TABLE public."Purchase" ALTER COLUMN "ID" DROP DEFAULT;
       public          postgres    false    212    213    213            Q          0    18385    Category 
   TABLE DATA                 public          postgres    false    214   ?V       S          0    18399    Comments 
   TABLE DATA                 public          postgres    false    216   ?V       ]          0    18658    DelUsers 
   TABLE DATA                 public          postgres    false    226   	W       Y          0    18430    MainDiscount 
   TABLE DATA                 public          postgres    false    222   ?W       T          0    18406    Manager 
   TABLE DATA                 public          postgres    false    217   ?W       ^          0    18665 	   Periphery 
   TABLE DATA                 public          postgres    false    227   ?X       \          0    18643    Posts 
   TABLE DATA                 public          postgres    false    225   ?Y       W          0    18421    PreOrder 
   TABLE DATA                 public          postgres    false    220   #[       R          0    18392    ProductCategory 
   TABLE DATA                 public          postgres    false    215   =[       M          0    18361    Products 
   TABLE DATA                 public          postgres    false    210   W[       P          0    18377    Purchase 
   TABLE DATA                 public          postgres    false    213   ?]       U          0    18413    Role 
   TABLE DATA                 public          postgres    false    218   ?_       Z          0    18438    Series 
   TABLE DATA                 public          postgres    false    223   ?_       N          0    18369    Users 
   TABLE DATA                 public          postgres    false    211   *`       j           0    0    MainDiscount_ID_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public."MainDiscount_ID_seq"', 1, false);
          public          postgres    false    221            k           0    0    Posts_Id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public."Posts_Id_seq"', 33, true);
          public          postgres    false    224            l           0    0    PreOrder_ID_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public."PreOrder_ID_seq"', 1, false);
          public          postgres    false    219            m           0    0    Products_Article_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public."Products_Article_seq"', 1, false);
          public          postgres    false    209            n           0    0    Purchase_ID_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public."Purchase_ID_seq"', 3, true);
          public          postgres    false    212            ?           2606    18391    Category Category_pk 
   CONSTRAINT     Z   ALTER TABLE ONLY public."Category"
    ADD CONSTRAINT "Category_pk" PRIMARY KEY ("Name");
 B   ALTER TABLE ONLY public."Category" DROP CONSTRAINT "Category_pk";
       public            postgres    false    214            ?           2606    18405    Comments Comments_pk 
   CONSTRAINT     X   ALTER TABLE ONLY public."Comments"
    ADD CONSTRAINT "Comments_pk" PRIMARY KEY ("ID");
 B   ALTER TABLE ONLY public."Comments" DROP CONSTRAINT "Comments_pk";
       public            postgres    false    216            ?           2606    18664    DelUsers DelUsers_pkey 
   CONSTRAINT     ]   ALTER TABLE ONLY public."DelUsers"
    ADD CONSTRAINT "DelUsers_pkey" PRIMARY KEY ("Email");
 D   ALTER TABLE ONLY public."DelUsers" DROP CONSTRAINT "DelUsers_pkey";
       public            postgres    false    226            ?           2606    18437    MainDiscount MainDiscount_pk 
   CONSTRAINT     `   ALTER TABLE ONLY public."MainDiscount"
    ADD CONSTRAINT "MainDiscount_pk" PRIMARY KEY ("ID");
 J   ALTER TABLE ONLY public."MainDiscount" DROP CONSTRAINT "MainDiscount_pk";
       public            postgres    false    222            ?           2606    18412    Manager Manager_pk 
   CONSTRAINT     Y   ALTER TABLE ONLY public."Manager"
    ADD CONSTRAINT "Manager_pk" PRIMARY KEY ("Email");
 @   ALTER TABLE ONLY public."Manager" DROP CONSTRAINT "Manager_pk";
       public            postgres    false    217            ?           2606    18671    Periphery Periphery_pk 
   CONSTRAINT     _   ALTER TABLE ONLY public."Periphery"
    ADD CONSTRAINT "Periphery_pk" PRIMARY KEY ("Article");
 D   ALTER TABLE ONLY public."Periphery" DROP CONSTRAINT "Periphery_pk";
       public            postgres    false    227            ?           2606    18650    Posts Post_pk 
   CONSTRAINT     Q   ALTER TABLE ONLY public."Posts"
    ADD CONSTRAINT "Post_pk" PRIMARY KEY ("Id");
 ;   ALTER TABLE ONLY public."Posts" DROP CONSTRAINT "Post_pk";
       public            postgres    false    225            ?           2606    18428    PreOrder PreOrder_pk 
   CONSTRAINT     X   ALTER TABLE ONLY public."PreOrder"
    ADD CONSTRAINT "PreOrder_pk" PRIMARY KEY ("ID");
 B   ALTER TABLE ONLY public."PreOrder" DROP CONSTRAINT "PreOrder_pk";
       public            postgres    false    220            ?           2606    18398 "   ProductCategory ProductCategory_pk 
   CONSTRAINT     w   ALTER TABLE ONLY public."ProductCategory"
    ADD CONSTRAINT "ProductCategory_pk" PRIMARY KEY ("Product", "Category");
 P   ALTER TABLE ONLY public."ProductCategory" DROP CONSTRAINT "ProductCategory_pk";
       public            postgres    false    215    215            ?           2606    18368    Products Products_pk 
   CONSTRAINT     ]   ALTER TABLE ONLY public."Products"
    ADD CONSTRAINT "Products_pk" PRIMARY KEY ("Article");
 B   ALTER TABLE ONLY public."Products" DROP CONSTRAINT "Products_pk";
       public            postgres    false    210            ?           2606    18384    Purchase Purchase_pk 
   CONSTRAINT     X   ALTER TABLE ONLY public."Purchase"
    ADD CONSTRAINT "Purchase_pk" PRIMARY KEY ("ID");
 B   ALTER TABLE ONLY public."Purchase" DROP CONSTRAINT "Purchase_pk";
       public            postgres    false    213            ?           2606    18419    Role Role_pk 
   CONSTRAINT     R   ALTER TABLE ONLY public."Role"
    ADD CONSTRAINT "Role_pk" PRIMARY KEY ("Name");
 :   ALTER TABLE ONLY public."Role" DROP CONSTRAINT "Role_pk";
       public            postgres    false    218            ?           2606    18444    Series Series_pk 
   CONSTRAINT     V   ALTER TABLE ONLY public."Series"
    ADD CONSTRAINT "Series_pk" PRIMARY KEY ("Name");
 >   ALTER TABLE ONLY public."Series" DROP CONSTRAINT "Series_pk";
       public            postgres    false    223            ?           2606    18375    Users Users_pk 
   CONSTRAINT     U   ALTER TABLE ONLY public."Users"
    ADD CONSTRAINT "Users_pk" PRIMARY KEY ("Email");
 <   ALTER TABLE ONLY public."Users" DROP CONSTRAINT "Users_pk";
       public            postgres    false    211            ?           2606    18470    Comments Comments_fk0    FK CONSTRAINT     ?   ALTER TABLE ONLY public."Comments"
    ADD CONSTRAINT "Comments_fk0" FOREIGN KEY ("Product") REFERENCES public."Products"("Article");
 C   ALTER TABLE ONLY public."Comments" DROP CONSTRAINT "Comments_fk0";
       public          postgres    false    3226    210    216            ?           2606    18475    Comments Comments_fk1    FK CONSTRAINT     ~   ALTER TABLE ONLY public."Comments"
    ADD CONSTRAINT "Comments_fk1" FOREIGN KEY ("User") REFERENCES public."Users"("Email");
 C   ALTER TABLE ONLY public."Comments" DROP CONSTRAINT "Comments_fk1";
       public          postgres    false    211    3228    216            ?           2606    18490    MainDiscount MainDiscount_fk0    FK CONSTRAINT     ?   ALTER TABLE ONLY public."MainDiscount"
    ADD CONSTRAINT "MainDiscount_fk0" FOREIGN KEY ("Product") REFERENCES public."Series"("Name");
 K   ALTER TABLE ONLY public."MainDiscount" DROP CONSTRAINT "MainDiscount_fk0";
       public          postgres    false    3246    223    222            ?           2606    18495    MainDiscount MainDiscount_fk1    FK CONSTRAINT     ?   ALTER TABLE ONLY public."MainDiscount"
    ADD CONSTRAINT "MainDiscount_fk1" FOREIGN KEY ("Name") REFERENCES public."Products"("Article");
 K   ALTER TABLE ONLY public."MainDiscount" DROP CONSTRAINT "MainDiscount_fk1";
       public          postgres    false    3226    210    222            ?           2606    18480    Manager Manager_fk0    FK CONSTRAINT     z   ALTER TABLE ONLY public."Manager"
    ADD CONSTRAINT "Manager_fk0" FOREIGN KEY ("Role") REFERENCES public."Role"("Name");
 A   ALTER TABLE ONLY public."Manager" DROP CONSTRAINT "Manager_fk0";
       public          postgres    false    3240    217    218            ?           2606    18672    Periphery Periphery_fk0    FK CONSTRAINT     ?   ALTER TABLE ONLY public."Periphery"
    ADD CONSTRAINT "Periphery_fk0" FOREIGN KEY ("Series") REFERENCES public."Series"("Name");
 E   ALTER TABLE ONLY public."Periphery" DROP CONSTRAINT "Periphery_fk0";
       public          postgres    false    3246    227    223            ?           2606    18485    PreOrder PreOrder_fk0    FK CONSTRAINT     ?   ALTER TABLE ONLY public."PreOrder"
    ADD CONSTRAINT "PreOrder_fk0" FOREIGN KEY ("Product") REFERENCES public."Series"("Name");
 C   ALTER TABLE ONLY public."PreOrder" DROP CONSTRAINT "PreOrder_fk0";
       public          postgres    false    223    220    3246            ?           2606    18460 #   ProductCategory ProductCategory_fk0    FK CONSTRAINT     ?   ALTER TABLE ONLY public."ProductCategory"
    ADD CONSTRAINT "ProductCategory_fk0" FOREIGN KEY ("Product") REFERENCES public."Products"("Article");
 Q   ALTER TABLE ONLY public."ProductCategory" DROP CONSTRAINT "ProductCategory_fk0";
       public          postgres    false    215    3226    210            ?           2606    18465 #   ProductCategory ProductCategory_fk1    FK CONSTRAINT     ?   ALTER TABLE ONLY public."ProductCategory"
    ADD CONSTRAINT "ProductCategory_fk1" FOREIGN KEY ("Category") REFERENCES public."Category"("Name");
 Q   ALTER TABLE ONLY public."ProductCategory" DROP CONSTRAINT "ProductCategory_fk1";
       public          postgres    false    3232    214    215            ?           2606    18445    Products Products_fk0    FK CONSTRAINT     ?   ALTER TABLE ONLY public."Products"
    ADD CONSTRAINT "Products_fk0" FOREIGN KEY ("Series") REFERENCES public."Series"("Name");
 C   ALTER TABLE ONLY public."Products" DROP CONSTRAINT "Products_fk0";
       public          postgres    false    223    3246    210            ?           2606    18450    Purchase Purchase_fk0    FK CONSTRAINT     ~   ALTER TABLE ONLY public."Purchase"
    ADD CONSTRAINT "Purchase_fk0" FOREIGN KEY ("User") REFERENCES public."Users"("Email");
 C   ALTER TABLE ONLY public."Purchase" DROP CONSTRAINT "Purchase_fk0";
       public          postgres    false    213    3228    211            ?           2606    18455    Purchase Purchase_fk1    FK CONSTRAINT     ?   ALTER TABLE ONLY public."Purchase"
    ADD CONSTRAINT "Purchase_fk1" FOREIGN KEY ("Product") REFERENCES public."Products"("Article");
 C   ALTER TABLE ONLY public."Purchase" DROP CONSTRAINT "Purchase_fk1";
       public          postgres    false    210    213    3226            Q   M   x???v
Q???W((M??L?SrN,IM?/?TRs?	uV?PH-?,?H-?T״??$NOpnbQIAF~^*H ? ?      S   
   x???          ]   ?   x???v
Q???W((M??L?SrI?	-N-*VRs?	uV?P?H???/?/?IqH?M???K??U?QPw?,*.Q(?/W??SHT(IL?IU״??$????T?.,??|a????Ծ???.l ??D????"#c4S?_?4??¾;.콰??. ????{H0? ==?? ?=???;:'????????????t?1\\ u˞)      Y   
   x???          T   ?   x??нj?0 ??O!?8?b\??B??%?C?B?f?Ȳ,?W????kuS'?????CU??ZT??;????$?\@?f?????ؠm
?w?G}?y?L?2}Di!????@'???j/`???%}?'U?7h?gM?%??R??ܶko?M????|??*?=KK?J???????ֆ????75n??5???O??I*??Y-fG???????َ?;???B???ݳ?M?__?G      ^   ?   x???O?@??~????????N??H??.??%?~?6C3???3?͏7???7؎??$ߟx?W?,?I??B?????]????@?9??\???T\?"_??A=9?C*??Tz????????Tw???1?XzlX?`?x
???ɚL?0????0c¨J9????M?5?e??`??L6??_ ?<?$?n&????<?F???>???S+?+????      \   3  x???QO?0?w??????lC??DpqA>??\?l?\[?Ƿ?ѐ@_zm.???6?<??l?W???fw?X3??";?9??i?QNJańUI??$?j???"͋?^?????x??(?3????nt?8?\?Dg???dX??P?/U?3J??'?5?5?A?k????h?BjeHZ???(?eÒU?نxbcYI?: ib????ȷ?*?s?N?!e?G(???G?Y|????tF4΀Տo?S????Q??vQ???ޯ?|?z????{?sʆW??{?tsə???b?N?/ا?!+?c?|??`????      W   
   x???          R   
   x???          M   q  x??V]O?0}ﯸ⥛4???8֞&??誶?`on⁷4??!ƿ?vRH)R???????{????b?????;??6???ɲ2Y?6?	??????>`??'?/Mj???:??>?H?E??Y&y??:5mќՍ??/???ӿ??s?6?
Ꝭ???j????*?:???W*?iw??4
X!?V?dR,???dc?`?֧A?/ח??EcJ???p?\96-??z>o?4]?_?lc??=1 Aِ?????p??,??JI?????@z?d?.0I?đ???a?'cb?'q?^???o<2?
[U7P?*?2????G?V?#?g??b?s?a????????N??$b???S#???afy??c?`??($?F29?uXH??ie????CV?t?JV?Lq??B???CE(?^ ?0:????<?$L???5??lsY???E??i,?S?1/&^??t???vs??p???s?? 	??????#>?<??wj?/??q}E>N_z??'~<?????_????/c?;??Awެ0\??)<Wy#{??u?K??zY?????P??P?l?|c߉?=?Fţ???Bvh_?ݵ.????$̷?>[
???????0????ž{g/:"?5?=??'?      P   ?  x????N?0??{???@???m]. ?a?{WU?Z???!x{?	6??p顶???mgt?pu?V????Z?'U??????9[G??????A?@???"{??Y?Ft>????M?($m?w?41??$ :E<?D!?????Yo?	G?ȴ???q??{h?7A??]??i[???????켜O???r??Z?m?b?]???fI?m??E?z?.??K??d????0???>?c[$?ꢺ???X?a?^"??u?nrJ?]XԺ(???ܦ`?x4??V)?I1??Y??K?a|?%.'
?9?????u?ӑ_?$D?????h?N	?N I?@l}T6@v"???$?7??b8 Hi/?05?]????<l??Pw?i?i???xg??Ф?m?I??	r?;-oxd???(???>z)?:      U   9   x???v
Q???W((M??L?S
??IURs?	uV?POL???S?Q0Դ??? mka      Z   G   x???v
Q???W((M??L?S
N-?L-VRs?	uV?P?O?W״??$FqPjJn&Ѫ#2?!ʹ? ~?(?      N   ?  x???mk?@???S???ôN??&IJ`??N{5dG?	߃w:??>?|N?p??1|????g?$??1]?l????)??LR_?-_E__v?4?0??y??d?5?o?ѭ?~??OÏ??V\?ARiɨ??y??|=?2???}?w??tvU#56?ԧտjuRM2???#hJ?'P????1??gY??J????%?!?=)??????a?8?R?lp?
R??E????H?ˋ??$?ay??fZv?Cy??T???E?,|w?J??W;|?9(?????	S[?%???rѵ??
t(???I?xv??J??-?j???SE???2????????p?aWX䟲Ol?? ???a??Z?ʁ????^Ϧ????(?????e? x?
?>}`?`?M7S?~?(????.?|o??_g??;`?j0??2?v     