USE [master]
GO
/****** Object:  Database [mygram_db]    Script Date: 8/5/2022 6:42:33 PM ******/
CREATE DATABASE [mygram_db]
 CONTAINMENT = NONE
 ON  PRIMARY 
( NAME = N'mygram_db', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL15.MSSQLSERVER\MSSQL\DATA\mygram_db.mdf' , SIZE = 8192KB , MAXSIZE = UNLIMITED, FILEGROWTH = 65536KB )
 LOG ON 
( NAME = N'mygram_db_log', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL15.MSSQLSERVER\MSSQL\DATA\mygram_db_log.ldf' , SIZE = 8192KB , MAXSIZE = 2048GB , FILEGROWTH = 65536KB )
 WITH CATALOG_COLLATION = DATABASE_DEFAULT
GO
ALTER DATABASE [mygram_db] SET COMPATIBILITY_LEVEL = 150
GO
IF (1 = FULLTEXTSERVICEPROPERTY('IsFullTextInstalled'))
begin
EXEC [mygram_db].[dbo].[sp_fulltext_database] @action = 'enable'
end
GO
ALTER DATABASE [mygram_db] SET ANSI_NULL_DEFAULT OFF 
GO
ALTER DATABASE [mygram_db] SET ANSI_NULLS OFF 
GO
ALTER DATABASE [mygram_db] SET ANSI_PADDING OFF 
GO
ALTER DATABASE [mygram_db] SET ANSI_WARNINGS OFF 
GO
ALTER DATABASE [mygram_db] SET ARITHABORT OFF 
GO
ALTER DATABASE [mygram_db] SET AUTO_CLOSE OFF 
GO
ALTER DATABASE [mygram_db] SET AUTO_SHRINK OFF 
GO
ALTER DATABASE [mygram_db] SET AUTO_UPDATE_STATISTICS ON 
GO
ALTER DATABASE [mygram_db] SET CURSOR_CLOSE_ON_COMMIT OFF 
GO
ALTER DATABASE [mygram_db] SET CURSOR_DEFAULT  GLOBAL 
GO
ALTER DATABASE [mygram_db] SET CONCAT_NULL_YIELDS_NULL OFF 
GO
ALTER DATABASE [mygram_db] SET NUMERIC_ROUNDABORT OFF 
GO
ALTER DATABASE [mygram_db] SET QUOTED_IDENTIFIER OFF 
GO
ALTER DATABASE [mygram_db] SET RECURSIVE_TRIGGERS OFF 
GO
ALTER DATABASE [mygram_db] SET  DISABLE_BROKER 
GO
ALTER DATABASE [mygram_db] SET AUTO_UPDATE_STATISTICS_ASYNC OFF 
GO
ALTER DATABASE [mygram_db] SET DATE_CORRELATION_OPTIMIZATION OFF 
GO
ALTER DATABASE [mygram_db] SET TRUSTWORTHY OFF 
GO
ALTER DATABASE [mygram_db] SET ALLOW_SNAPSHOT_ISOLATION OFF 
GO
ALTER DATABASE [mygram_db] SET PARAMETERIZATION SIMPLE 
GO
ALTER DATABASE [mygram_db] SET READ_COMMITTED_SNAPSHOT OFF 
GO
ALTER DATABASE [mygram_db] SET HONOR_BROKER_PRIORITY OFF 
GO
ALTER DATABASE [mygram_db] SET RECOVERY FULL 
GO
ALTER DATABASE [mygram_db] SET  MULTI_USER 
GO
ALTER DATABASE [mygram_db] SET PAGE_VERIFY CHECKSUM  
GO
ALTER DATABASE [mygram_db] SET DB_CHAINING OFF 
GO
ALTER DATABASE [mygram_db] SET FILESTREAM( NON_TRANSACTED_ACCESS = OFF ) 
GO
ALTER DATABASE [mygram_db] SET TARGET_RECOVERY_TIME = 60 SECONDS 
GO
ALTER DATABASE [mygram_db] SET DELAYED_DURABILITY = DISABLED 
GO
ALTER DATABASE [mygram_db] SET ACCELERATED_DATABASE_RECOVERY = OFF  
GO
EXEC sys.sp_db_vardecimal_storage_format N'mygram_db', N'ON'
GO
ALTER DATABASE [mygram_db] SET QUERY_STORE = OFF
GO
USE [mygram_db]
GO
/****** Object:  Table [dbo].[comments]    Script Date: 8/5/2022 6:42:34 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[comments](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[message] [nvarchar](max) NOT NULL,
	[created_at] [datetime] NOT NULL,
	[updated_at] [datetime] NOT NULL,
	[user_id] [int] NOT NULL,
	[photo_id] [int] NOT NULL,
 CONSTRAINT [PK_comments] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[photos]    Script Date: 8/5/2022 6:42:34 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[photos](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[title] [nvarchar](100) NOT NULL,
	[caption] [nvarchar](100) NULL,
	[photo_url] [nvarchar](200) NOT NULL,
	[created_at] [datetime] NOT NULL,
	[updated_at] [datetime] NOT NULL,
	[user_id] [int] NOT NULL,
 CONSTRAINT [PK_photos] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[social_medias]    Script Date: 8/5/2022 6:42:34 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[social_medias](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[name] [nvarchar](100) NOT NULL,
	[social_media_url] [nvarchar](200) NOT NULL,
	[user_id] [int] NOT NULL,
	[created_at] [datetime] NOT NULL,
	[updated_at] [datetime] NOT NULL,
 CONSTRAINT [PK_social_medias] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[users]    Script Date: 8/5/2022 6:42:34 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[users](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[username] [nvarchar](50) NOT NULL,
	[password] [nvarchar](100) NOT NULL,
	[email] [nvarchar](100) NOT NULL,
	[age] [int] NOT NULL,
	[created_at] [datetime] NOT NULL,
	[updated_at] [datetime] NOT NULL,
	[profile_image_url] [nvarchar](200) NOT NULL,
 CONSTRAINT [PK_users] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [IX_email_users]    Script Date: 8/5/2022 6:42:34 PM ******/
CREATE UNIQUE NONCLUSTERED INDEX [IX_email_users] ON [dbo].[users]
(
	[email] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, DROP_EXISTING = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [IX_username_users]    Script Date: 8/5/2022 6:42:34 PM ******/
CREATE UNIQUE NONCLUSTERED INDEX [IX_username_users] ON [dbo].[users]
(
	[username] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, DROP_EXISTING = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
ALTER TABLE [dbo].[comments] ADD  CONSTRAINT [DF_comments_created_at]  DEFAULT ((0)) FOR [created_at]
GO
ALTER TABLE [dbo].[comments] ADD  CONSTRAINT [DF_comments_updated_at]  DEFAULT ((0)) FOR [updated_at]
GO
ALTER TABLE [dbo].[photos] ADD  CONSTRAINT [DF_photos_created_at]  DEFAULT ((0)) FOR [created_at]
GO
ALTER TABLE [dbo].[photos] ADD  CONSTRAINT [DF_photos_updated_at]  DEFAULT ((0)) FOR [updated_at]
GO
ALTER TABLE [dbo].[social_medias] ADD  CONSTRAINT [DF_social_medias_created_at]  DEFAULT ((0)) FOR [created_at]
GO
ALTER TABLE [dbo].[social_medias] ADD  CONSTRAINT [DF_social_medias_updated_at]  DEFAULT ((0)) FOR [updated_at]
GO
ALTER TABLE [dbo].[users] ADD  CONSTRAINT [DF_users_created_at]  DEFAULT ((0)) FOR [created_at]
GO
ALTER TABLE [dbo].[users] ADD  CONSTRAINT [DF_users_updated_at]  DEFAULT ((0)) FOR [updated_at]
GO
ALTER TABLE [dbo].[comments]  WITH CHECK ADD  CONSTRAINT [FK_comments_photos] FOREIGN KEY([photo_id])
REFERENCES [dbo].[photos] ([id])
GO
ALTER TABLE [dbo].[comments] CHECK CONSTRAINT [FK_comments_photos]
GO
ALTER TABLE [dbo].[comments]  WITH CHECK ADD  CONSTRAINT [FK_comments_users] FOREIGN KEY([user_id])
REFERENCES [dbo].[users] ([id])
GO
ALTER TABLE [dbo].[comments] CHECK CONSTRAINT [FK_comments_users]
GO
ALTER TABLE [dbo].[photos]  WITH CHECK ADD  CONSTRAINT [FK_photos_users] FOREIGN KEY([user_id])
REFERENCES [dbo].[users] ([id])
GO
ALTER TABLE [dbo].[photos] CHECK CONSTRAINT [FK_photos_users]
GO
ALTER TABLE [dbo].[social_medias]  WITH CHECK ADD  CONSTRAINT [FK_social_medias_users] FOREIGN KEY([user_id])
REFERENCES [dbo].[users] ([id])
GO
ALTER TABLE [dbo].[social_medias] CHECK CONSTRAINT [FK_social_medias_users]
GO
USE [master]
GO
ALTER DATABASE [mygram_db] SET  READ_WRITE 
GO