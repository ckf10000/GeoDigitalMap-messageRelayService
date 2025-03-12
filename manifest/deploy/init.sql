-- 创建数据库
create database geo_digital_map owner admin;

create table if not exists "t_msg_event_how" (
	id serial primary key,
	code varchar(50) not null,
	how varchar(50) not null,
    alias varchar(100) not null default '',
	description text not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    deleted_at bigint default 0,
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_how is '消息事件触发方式配置表';
comment on column t_msg_event_how.id is '自增主键，唯一标识记录';
comment on column t_msg_event_how.code is '事件类型编码(长度限制50字符)，不可为空';
comment on column t_msg_event_how.how is '事件触发方式(长度限制50字符)，不可为空';
comment on column t_msg_event_how.alias is '事件别名(长度限制100字符)，默认空字符串';
comment on column t_msg_event_how.description is '事件详细描述，默认空文本';
comment on column t_msg_event_how.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_how.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_how.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_how.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_how.deleted_at is '软删时间(unix时间戳，秒级)，默认0表示未删除';
comment on column t_msg_event_how.is_deleted is '软删标志，默认false表示有效数据';

create unique index idx_unique_msg_event_how_deleted_at on t_msg_event_how(id, deleted_at);
create index idx_msg_event_how_alias on t_msg_event_how(alias);

-- 创建触发器
create trigger update_msg_event_how_updated_at
before update on t_msg_event_how
for each row
execute function update_timestamp();

insert into t_msg_event_how ("code", "how", "alias") values ('m-i', 'mensurated', '测量');
insert into t_msg_event_how ("code", "how", "alias") values ('h-t', 'human', '人类');
insert into t_msg_event_how ("code", "how", "alias") values ('h-t', 'retyped', '重新输入');
insert into t_msg_event_how ("code", "how", "alias") values ('m-', 'machine', '机器');
insert into t_msg_event_how ("code", "how", "alias") values ('m-g', 'gps', '全球定位系统');
insert into t_msg_event_how ("code", "how", "alias") values ('h-g-i-g-o', 'nonCoT', '非协同作战');
insert into t_msg_event_how ("code", "how", "alias") values ('h-g-i-g-o', 'gigo', '无用输入输出');
insert into t_msg_event_how ("code", "how", "alias") values ('a-f-G-E-V-9-1-1', 'mayday', '呼救信号');
insert into t_msg_event_how ("code", "how", "alias") values ('h-e', 'estimated', '估计的');
insert into t_msg_event_how ("code", "how", "alias") values ('h-c', 'calculated', '计算的');
insert into t_msg_event_how ("code", "how", "alias") values ('h-t', 'transcribed', '转录的');
insert into t_msg_event_how ("code", "how", "alias") values ('h-p', 'pasted', '粘贴的');
insert into t_msg_event_how ("code", "how", "alias") values ('m-m', 'magnetic', '有磁性的');
insert into t_msg_event_how ("code", "how", "alias") values ('m-n', 'ins', '惯性导航系统');
insert into t_msg_event_how ("code", "how", "alias") values ('m-s', 'simulated', '模拟的');
insert into t_msg_event_how ("code", "how", "alias") values ('m-c', 'configured', '配置的');
insert into t_msg_event_how ("code", "how", "alias") values ('m-r', 'radio', '无线电');
insert into t_msg_event_how ("code", "how", "alias") values ('m-p', 'passed', '通过的');
insert into t_msg_event_how ("code", "how", "alias") values ('m-p', 'propagated', '测量');
insert into t_msg_event_how ("code", "how", "alias") values ('m-f', 'fused', '融合的');
insert into t_msg_event_how ("code", "how", "alias") values ('m-a', 'tracker', '追踪的');
insert into t_msg_event_how ("code", "how", "alias") values ('m-g-n', 'ins+gps', '惯性导航系统加全球定位系统');
insert into t_msg_event_how ("code", "how", "alias") values ('m-g-d', 'dgps', '差分全球定位系统');
insert into t_msg_event_how ("code", "how", "alias") values ('m-r-e', 'eplrs', '增强型位置报告系统');
insert into t_msg_event_how ("code", "how", "alias") values ('m-r-p', 'plrs', '位置报告系统');
insert into t_msg_event_how ("code", "how", "alias") values ('m-r-d', 'doppler', '多普勒');
insert into t_msg_event_how ("code", "how", "alias") values ('m-r-v', 'vhf', '甚高频');
insert into t_msg_event_how ("code", "how", "alias") values ('m-r-t', 'tadil', '战术数字信息链路');
insert into t_msg_event_how ("code", "how", "alias") values ('m-r-t-a', 'tadila', '战术数字信息链路A');
insert into t_msg_event_how ("code", "how", "alias") values ('m-r-t-b', 'tadilb', '战术数字信息链路B');
insert into t_msg_event_how ("code", "how", "alias") values ('m-r-t-j', 'tadilj', '战术数字信息链路J');


create table if not exists "t_msg_event_type" (
	id serial primary key,
	event_type varchar(50) not null,
    alias varchar(100) not null default '',
	description text not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    deleted_at bigint default 0,
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_type is '消息事件类型配置表';
comment on column t_msg_event_type.id is '自增主键，唯一标识记录';
comment on column t_msg_event_type.event_type is '事件类型编码(长度限制50字符)，不可为空';
comment on column t_msg_event_type.alias is '事件类型别名(长度限制100字符)，默认空字符串';
comment on column t_msg_event_type.description is '事件类型详细描述，默认空文本';
comment on column t_msg_event_type.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_type.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_type.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_type.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_type.deleted_at is '软删时间(unix时间戳，秒级)，默认0表示未删除';
comment on column t_msg_event_type.is_deleted is '软删标志，默认false表示有效数据';

create unique index idx_unique_msg_event_type_deleted_at on t_msg_event_type(id, deleted_at);
create index idx_msg_event_type_alias on t_msg_event_type(alias);

-- 创建触发器
create trigger update_msg_event_type_updated_at
before update on t_msg_event_type
for each row
execute function update_timestamp();

insert into t_msg_event_type ("event_type", "alias") values ('t-x-c-t', 'Ping');
insert into t_msg_event_type ("event_type", "alias") values ('t-x-c-t-r', 'TakPong');
insert into t_msg_event_type ("event_type", "alias") values ('b-t-f', 'GeoChat');
insert into t_msg_event_type ("event_type", "alias") values ('a-f-G-U', 'UserUpdate');
insert into t_msg_event_type ("event_type", "alias") values ('a-f-G-U-C', 'UserUpdate');
insert into t_msg_event_type ("event_type", "alias") values ('a-f-G-U-C-I', 'UserUpdate');
insert into t_msg_event_type ("event_type", "alias") values ('a-f-G-E-V', 'UserUpdate');
insert into t_msg_event_type ("event_type", "alias") values ('a-f-G-E-V-A', 'UserUpdate');
insert into t_msg_event_type ("event_type", "alias") values ('a-f-G-E-V-C', 'UserUpdate');
insert into t_msg_event_type ("event_type", "alias") values ('a-h-G', 'Point');
insert into t_msg_event_type ("event_type", "alias") values ('a-n-G', 'Point');
insert into t_msg_event_type ("event_type", "alias") values ('a-f-G', 'Point');
insert into t_msg_event_type ("event_type", "alias") values ('a-u-G', 'Point');
insert into t_msg_event_type ("event_type", "alias") values ('t-x-m-c', 'Xml');
insert into t_msg_event_type ("event_type", "alias") values ('u-r-b-bullseye', 'bullseye');


-- 主表创建（需优先执行）
create table if not exists "t_msg_event" (
	id varchar(100) primary key,
    xml_string varchar(1000) not null default '',
	sender varchar(100) not null,
    receiver varchar(100) not null default '',
	send_result int not null default 1,
    occurrence_time bigint not null default 0,
    checksum varchar(32) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event is '消息事件记录收发及状态信息表-主表';
comment on column t_msg_event.id is '事件唯一标识符(主键)';
comment on column t_msg_event.xml_string is '事件xml原始数据(长度限制1000字符)，默认空字符串';
comment on column t_msg_event.sender is '事件发送方标识符(长度限制100字符)，不可为空';
comment on column t_msg_event.receiver is '事件接收方标识符(长度限制100字符)，默认空字符串';
comment on column t_msg_event.send_result is '发送结果状态码(1-成功，0-失败)，默认1';
comment on column t_msg_event.occurrence_time is '事件发生时间(unix时间戳，秒级)，默认0表示未记录';
comment on column t_msg_event.checksum is '数据校验和(md5哈希值)，默认空字符串';
comment on column t_msg_event.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event.is_deleted is '软删标志，默认false表示有效数据';

create unique index idx_unique_msg_event_id on t_msg_event(id);
create index idx_msg_event_xml_string on t_msg_event(xml_string);
create index idx_msg_event_sender on t_msg_event(sender);
create index idx_msg_event_receiver on t_msg_event(receiver);
create index idx_msg_event_send_result on t_msg_event(send_result);

-- 创建触发器
create trigger update_msg_event_updated_at
before update on t_msg_event
for each row
execute function update_timestamp();

-- 一级关联表
create table if not exists "t_msg_event_summary" (
    event_id varchar(100) primary key,
    how varchar(100) not null default '',
	time bigint not null default 0,
    start bigint not null default 0,
    stale bigint not null default 0,
    type varchar(100) not null default '',
    version varchar(100) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_summary is '消息事件生命周期摘要表-一级关联表';
comment on column t_msg_event_summary.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_summary.how is '事件触发方式标识(长度限制100字符)，默认空字符串';
comment on column t_msg_event_summary.time is '事件发生时间(unix时间戳，秒级)，默认0表示未记录';
comment on column t_msg_event_summary.start is '事件有效开始时间(unix时间戳，秒级)，默认0表示立即生效';
comment on column t_msg_event_summary.stale is '事件失效时间(unix时间戳，秒级)，默认0表示永不过期';
comment on column t_msg_event_summary.type is '事件分类类型(长度限制100字符)，默认空字符串';
comment on column t_msg_event_summary.version is '事件版本标识符(长度限制100字符)，默认空字符串';
comment on column t_msg_event_summary.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_summary.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_summary.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_summary.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_summary.is_deleted is '软删标志，默认false表示有效数据';

create unique index idx_unique_msg_event_summary_event_id on t_msg_event_summary(event_id);

-- 创建触发器
create trigger update_msg_event_summary_updated_at
before update on t_msg_event_summary
for each row
execute function update_timestamp();

-- 一级关联表
create table if not exists "t_msg_event_point" (
    event_id varchar(100) primary key,
    lat varchar(100) not null default '0',
    lon varchar(100) not null default '0',
	hae varchar(100) not null default '9999999.0',
    ce varchar(100) not null default '9999999.0',
    le varchar(100) not null default '9999999.0',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_point is '消息记录事件发生位置及定位精度元数据表-一级关联表';
comment on column t_msg_event_point.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_point.lat is '纬度坐标(wgs84坐标系)，默认0表示未记录，单位：度';
comment on column t_msg_event_point.lon is '经度坐标(gs84坐标系)，默认0表示未记录，单位：度';
comment on column t_msg_event_point.hae is '海拔高度(height above ellipsoid)，默认9999999.0表示无效值，单位：米';
comment on column t_msg_event_point.ce is '圆误差概率(circular error)，默认9999999.0表示无效值，单位：米';
comment on column t_msg_event_point.le is '线误差概率(linear error)，默认9999999.0表示无效值，单位：米';
comment on column t_msg_event_point.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_point.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_point.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_point.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_point.is_deleted is '软删标志，默认false表示有效数据';

create unique index idx_unique_msg_event_point_event_id on t_msg_event_point(event_id);

-- 创建触发器
create trigger update_msg_event_point_updated_at
before update on t_msg_event_point
for each row
execute function update_timestamp();

-- 一级关联表
create table if not exists "t_msg_event_sensor" (
    event_id varchar(100) primary key,
    elevation varchar(100) not null default '',
    vfov varchar(100) not null default '',
    north varchar(100) not null default '',
    roll varchar(100) not null default '',
    range varchar(100) not null default '',
    azimuth varchar(100) not null default '',
    model varchar(100) not null default '',
    fov varchar(100) not null default '',
    type varchar(100) not null default '',
    version varchar(100) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);

comment on table t_msg_event_sensor is '消息事件传感器参数表-一级关联表';
comment on column t_msg_event_sensor.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_sensor.elevation is '传感器海拔高度(单位：米)，默认空字符串';
comment on column t_msg_event_sensor.vfov is '垂直视场角(vertical fov)，默认空字符串';
comment on column t_msg_event_sensor.north is '北向校准参数，默认空字符串';
comment on column t_msg_event_sensor.roll is '传感器滚动角参数，默认空字符串';
comment on column t_msg_event_sensor.range is '传感器监测范围，默认空字符串';
comment on column t_msg_event_sensor.azimuth is '传感器方位角参数，默认空字符串';
comment on column t_msg_event_sensor.model is '传感器型号，默认空字符串';
comment on column t_msg_event_sensor.fov is '传感器视场角，默认空字符串';
comment on column t_msg_event_sensor.type is '传感器类型，默认空字符串';
comment on column t_msg_event_sensor.version is '传感器版本号，默认空字符串';
comment on column t_msg_event_sensor.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_sensor.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_sensor.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_sensor.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_sensor.is_deleted is '软删标志，默认false表示有效数据';

create unique index idx_unique_msg_event_sensor_event_id on t_msg_event_sensor(event_id);

-- 创建触发器
create trigger update_msg_event_sensor_updated_at
before update on t_msg_event_sensor
for each row
execute function update_timestamp();

-- 一级关联表
create table if not exists "t_msg_event_video" (
    event_id varchar(100) primary key,
    sensor varchar(100) not null default '',
    spi varchar(100) not null default '',
    url varchar(100) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);

comment on table t_msg_event_video is '消息事件视频元数据表-一级关联表';
comment on column t_msg_event_video.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_video.sensor is '视频传感器标识符(长度限制100字符)，默认空字符串';
comment on column t_msg_event_video.spi is '安全参数索引(security parameters index，长度限制100字符)，默认空字符串';
comment on column t_msg_event_video.url is '视频资源访问地址(长度限制100字符)，默认空字符串';
comment on column t_msg_event_video.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_video.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_video.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_video.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_video.is_deleted is '软删标志，默认false表示有效数据';


create unique index idx_unique_msg_event_video_event_id on t_msg_event_video(event_id);

-- 创建触发器
create trigger update_msg_event_video_updated_at
before update on t_msg_event_video
for each row
execute function update_timestamp();

-- 一级关联表
create table if not exists "t_msg_event_usericon" (
    event_id varchar(100) primary key,
    iconsetpath varchar(100) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);

comment on table t_msg_event_usericon is '消息事件用户图标配置表-一级关联表';
comment on column t_msg_event_usericon.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_usericon.iconsetpath is '用户图标集存储路径(相对路径，长度限制100字符)，默认空字符串';
comment on column t_msg_event_usericon.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_usericon.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_usericon.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_usericon.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_usericon.is_deleted is '软删标志，默认false表示有效数据';

create unique index idx_unique_msg_event_usericon_event_id on t_msg_event_usericon(event_id);

-- 创建触发器
create trigger update_msg_event_usericon_updated_at
before update on t_msg_event_usericon
for each row
execute function update_timestamp();

-- 一级关联表
create table if not exists "t_msg_event_device_status" (
    event_id varchar(100) primary key,
    battery varchar(100) not null default '',
    readiness varchar(100) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_device_status is '消息事件的设备状态信息表-一级关联表';
comment on column t_msg_event_device_status.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_device_status.battery is '设备电量百分比(长度限制100字符)，默认空字符串';
comment on column t_msg_event_device_status.readiness is '设备就绪状态标识(长度限制100字符)，默认空字符串';
comment on column t_msg_event_device_status.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_device_status.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_device_status.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_device_status.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_device_status.is_deleted is '软删标志，默认false表示有效数据';

create unique index idx_unique_msg_event_device_status_event_id on t_msg_event_device_status(event_id);

-- 创建触发器
create trigger update_msg_event_device_status_updated_at
before update on t_msg_event_device_status
for each row
execute function update_timestamp();

-- 一级关联表
create table if not exists "t_msg_event_mcv" (
    event_id varchar(100) primary key,
    device varchar(100) not null default '',
    os varchar(100) not null default '',
    platform varchar(100) not null default '',
    version varchar(100) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_mcv is '消息事件客户端版本信息表-一级关联表';
comment on column t_msg_event_mcv.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_mcv.device is '物理设备型号(长度限制100字符)，默认空字符串';
comment on column t_msg_event_mcv.os is '设备操作系统类型(长度限制100字符)，默认空字符串';
comment on column t_msg_event_mcv.platform is '消息事件客户端平台变体名称(长度限制100字符)，默认空字符串';
comment on column t_msg_event_mcv.version is '消息事件客户端软件版本号(长度限制100字符)，默认空字符串';
comment on column t_msg_event_mcv.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_mcv.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_mcv.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_mcv.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_mcv.is_deleted is '软删标志，默认false表示有效数据';

create unique index idx_unique_msg_event_mcv_event_id on t_msg_event_mcv(event_id);

-- 创建触发器
create trigger update_msg_event_mcv_updated_at
before update on t_msg_event_mcv
for each row
execute function update_timestamp();

-- 一级关联表
create table if not exists "t_msg_event_track" (
    event_id varchar(100) primary key,
    course varchar(100) not null default '',
    ecourse varchar(100) not null default '',
    eslope varchar(100) not null default '',
    espeed varchar(100) not null default '',
    speed varchar(100) not null default '',
    version varchar(100) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_track is '消息事件移动轨迹数据表-一级关联表';
comment on column t_msg_event_track.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_track.course is '移动方向角度(单位：度，长度限制100字符)，默认空字符串';
comment on column t_msg_event_track.ecourse is '方向角测量误差(1σ标准差，长度限制100字符)，默认空字符串';
comment on column t_msg_event_track.eslope is '坡度测量误差(1σ标准差，长度限制100字符)，默认空字符串';
comment on column t_msg_event_track.espeed is '速度测量误差(1σ标准差，长度限制100字符)，默认空字符串';
comment on column t_msg_event_track.speed is '移动速度值(单位：米/秒，长度限制100字符)，默认空字符串';
comment on column t_msg_event_track.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_track.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_track.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_track.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_track.is_deleted is '软删标志，默认false表示有效数据';


create unique index idx_unique_msg_event_track_event_id on t_msg_event_track(event_id);

-- 创建触发器
create trigger update_msg_event_track_updated_at
before update on t_msg_event_track
for each row
execute function update_timestamp();

-- 一级关联表
create table if not exists "t_msg_event_uid" (
    event_id varchar(100) primary key,
    droid varchar(100) not null default '',
    version varchar(100) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_uid is '消息用户标识扩展信息表-一级关联表';
comment on column t_msg_event_uid.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_uid.droid is 'android设备标识符(长度限制100字符)，默认空字符串';
comment on column t_msg_event_uid.version is 'uid版本标识(长度限制100字符)，默认空字符串';
comment on column t_msg_event_uid.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_uid.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_uid.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_uid.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_uid.is_deleted is '软删标志，默认false表示有效数据';


create unique index idx_unique_msg_event_uid_event_id on t_msg_event_uid(event_id);

-- 创建触发器
create trigger update_msg_event_uid_updated_at
before update on t_msg_event_uid
for each row
execute function update_timestamp();


-- 一级关联表
create table if not exists "t_msg_event_server_destination" (
    event_id varchar(100) primary key,
    destinations varchar(100) not null,
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_server_destination is '消息事件服务器目标配置表-一级关联表';
comment on column t_msg_event_server_destination.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_server_destination.destinations is '目标地址组合(格式：ip:端口:协议:设备id)，不可为空';
comment on column t_msg_event_server_destination.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_server_destination.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_server_destination.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_server_destination.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_server_destination.is_deleted is '软删标志，默认false表示有效数据';

create unique index idx_unique_msg_event_server_destination_event_id on t_msg_event_server_destination(event_id);

-- 创建触发器
create trigger update_msg_event_server_destination_updated_at
before update on t_msg_event_server_destination
for each row
execute function update_timestamp();

-- 一级关联表
create table if not exists "t_msg_event_remarks" (
    event_id varchar(100) primary key,
    keywords varchar(100) not null default '',
    source varchar(100) not null default '',
    time bigint not null default 0,
    to_field varchar(100) not null default '',
    version varchar(100) not null default '',
    intag varchar(1000) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_remarks is '消息事件备注信息表-一级关联表';
comment on column t_msg_event_remarks.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_remarks.keywords is '逗号分隔的备注关键词(如"任务a,紧急")，默认空字符串';
comment on column t_msg_event_remarks.source is '备注来源标识符(发送方uid)，默认空字符串';
comment on column t_msg_event_remarks.time is '备注生成时间(unix时间戳，秒级)，默认0表示未记录';
comment on column t_msg_event_remarks.to_field is '目标接收方uid，默认空字符串表示广播';
comment on column t_msg_event_remarks.version is '备注格式版本标识，默认空字符串';
comment on column t_msg_event_remarks.intag is '扩展标签信息(最大长度1000字符)，默认空字符串';
comment on column t_msg_event_remarks.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_remarks.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_remarks.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_remarks.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_remarks.is_deleted is '软删标志，默认false表示有效数据';


create unique index idx_unique_msg_event_remarks_event_id on t_msg_event_remarks(event_id);

-- 创建触发器
create trigger update_msg_event_remarks_updated_at
before update on t_msg_event_remarks
for each row
execute function update_timestamp();


-- 精确定位表
create table if not exists "t_msg_event_precision_location" (
    event_id varchar(100) primary key,
    alt_src varchar(100) not null default '',
    geo_point_src varchar(100) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_precision_location is '消息事件记录坐标来源信息-一级关联表';
comment on column t_msg_event_precision_location.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_precision_location.alt_src is '高度数据源标识(如dted0)，默认空字符串表示未指定';
comment on column t_msg_event_precision_location.geo_point_src is '地理坐标数据源标识(如gps/北斗)，默认空字符串';
comment on column t_msg_event_precision_location.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_precision_location.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_precision_location.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_precision_location.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_precision_location.is_deleted is '软删标志，默认false表示有效数据';

create unique index idx_unique_msg_event_precision_location_event_id on t_msg_event_precision_location(event_id);

-- 创建触发器
create trigger update_msg_event_precision_location_updated_at
before update on t_msg_event_precision_location
for each row
execute function update_timestamp();



-- 链接信息表
create table if not exists "t_msg_event_link" (
    event_id varchar(100) primary key,
    mime varchar(100) not null default '',
    parent_callsign varchar(100) not null default '',
    production_time bigint not null default 0,
    relation varchar(100) not null default '',
    remarks varchar(100) not null default '',
    type varchar(100) not null default '',
    uid varchar(100) not null default '',
    url varchar(100) not null default '',
    version varchar(100) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
-- 表级注释建议
comment on table t_msg_event_link is '消息事件关联的外部资源链接信息-一级关联表';
comment on column t_msg_event_link.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_link.mime is '资源mime类型(rfc 3023标准，如 application/xml/text/xml)，默认空字符串';
comment on column t_msg_event_link.parent_callsign is '上级节点呼号标识(用于层级关系追踪)，默认空字符串';
comment on column t_msg_event_link.production_time is '资源生成时间(unix时间戳，秒级)，默认0表示未记录';
comment on column t_msg_event_link.relation is '关联关系类型(cot标准：subject/object/indirect-object)，默认空字符串';
comment on column t_msg_event_link.remarks is '链接附加说明(自由文本，最大长度100字符)，默认空字符串';
comment on column t_msg_event_link.type is '链接资源分类(如影像/文档/传感器数据)，默认空字符串';
comment on column t_msg_event_link.uid is '关联资源全局唯一标识符(遵循cot uid规范)，默认空字符串';
comment on column t_msg_event_link.url is '资源访问地址(url格式，支持短链接)，默认空字符串';
comment on column t_msg_event_link.version is '链接模式版本标识(格式：主版本.次版本)，默认空字符串';
comment on column t_msg_event_link.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_link.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_link.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_link.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_link.is_deleted is '软删标志，默认false表示有效数据';

create unique index idx_unique_msg_event_link_event_id on t_msg_event_link(event_id);

-- 创建触发器
create trigger update_msg_event_link_updated_at
before update on t_msg_event_link
for each row
execute function update_timestamp();


-- 应急信息表
create table if not exists "t_msg_event_emergency" (
    event_id varchar(100) primary key,
    alert varchar(100) not null default '',
    cancel varchar(100) not null default 'false' check ("cancel" in ('true','false')),
    type varchar(100) not null default 'constructor',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_emergency is '紧急事件相关状态信息-一级关联表';
comment on column t_msg_event_emergency.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_emergency.alert is '应急警报等级或类型(如"critical"/"warning")，默认空字符串';
comment on column t_msg_event_emergency.cancel is '应急信标取消状态(true=已取消，false=有效)，默认false';
comment on column t_msg_event_emergency.type is '应急事件分类类型(如医疗/消防/警务)，默认"constructor"';
comment on column t_msg_event_emergency.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_emergency.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_emergency.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_emergency.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_emergency.is_deleted is '软删标志，默认false表示有效数据';


create unique index idx_unique_msg_event_emergency_event_id on t_msg_event_emergency(event_id);

-- 创建触发器
create trigger update_msg_event_emergency_updated_at
before update on t_msg_event_emergency
for each row
execute function update_timestamp();

-- 联系人信息表
create table if not exists "t_msg_event_contact" (
    event_id varchar(100) primary key,
    callsign varchar(100) not null default '',
    dsn varchar(100) not null default '',
    email varchar(100) not null default '' check (email ~* '^[a-za-z0-9._%+-]+@[a-za-z0-9.-]+\.[a-za-z]{2,}$'),
    endpoint varchar(100) not null default '',
    freq varchar(100) not null default '',
    hostname varchar(100) not null default '',
    iconsetpath varchar(100) not null default '',
    modulation varchar(100) check ("modulation" in ('am','fm')),
    phone varchar(20) not null default '',
    version varchar(100) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_contact is '消息事件联系人信息表-一级关联表';
comment on column t_msg_event_contact.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_contact.callsign is '通信呼号标识(如"alpha-01")，默认空字符串';
comment on column t_msg_event_contact.dsn is '国防交换网络编号(defense switched network)，默认空字符串';
comment on column t_msg_event_contact.email is '电子邮箱地址(需符合name@domain格式)，默认空字符串';
comment on column t_msg_event_contact.endpoint is 'api终端地址(ip:port格式)，默认空字符串';
comment on column t_msg_event_contact.freq is '通信频率(单位：mhz，如"121.500")，默认空字符串';
comment on column t_msg_event_contact.hostname is '网络主机名(需dns可解析)，默认空字符串';
comment on column t_msg_event_contact.iconsetpath is '图标资源存储路径(相对路径)，默认空字符串';
comment on column t_msg_event_contact.modulation is '无线电调制方式(am=调幅/fm=调频)';
comment on column t_msg_event_contact.phone is '联系电话号码(国际格式如+86-13912345678)，默认空字符串';
comment on column t_msg_event_contact.version is '联系人模式版本号(格式：v1.2.3)，默认空字符串';
comment on column t_msg_event_contact.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_contact.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_contact.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_contact.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_contact.is_deleted is '软删标志，默认false表示有效数据';


create unique index idx_unique_msg_event_contact_event_id on t_msg_event_contact(event_id);

-- 创建触发器
create trigger update_msg_event_contact_updated_at
before update on t_msg_event_contact
for each row
execute function update_timestamp();

-- 颜色配置表
create table if not exists "t_msg_event_color" (
    event_id varchar(100) primary key,
    argb varchar(100) not null default '' check ("argb" ~ '^#[0-9a-fa-f]{6,8}$'),
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_color is '消息事件颜色配置表-一级关联表';
comment on column t_msg_event_color.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_color.argb is 'argb颜色编码(格式：#aarrggbb 或 #rrggbb)，如#ff00ff00表示不透明品红色';
comment on column t_msg_event_color.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_color.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_color.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_color.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_color.is_deleted is '软删标志，默认false表示有效数据';


create unique index idx_unique_msg_event_color_event_id on t_msg_event_color(event_id);

-- 创建触发器
create trigger update_msg_event_color_updated_at
before update on t_msg_event_color
for each row
execute function update_timestamp();

-- 聊天信息表
create table if not exists "t_msg_event_chat" (
    event_id varchar(100) primary key,
    chat_room varchar(100) not null default '',
    group_owner varchar(100) not null default '',
    id varchar(100) not null default '',
    parent varchar(100) not null default '',
    sender_call_sign varchar(100) not null default '',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_chat is '消息事件聊天记录表-一级关联表';
comment on column t_msg_event_chat.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_chat.chat_room is '聊天室唯一标识符(如"cmd-center-01")，默认空字符串';
comment on column t_msg_event_chat.group_owner is '群组管理员标识符(uid格式)，默认空字符串';
comment on column t_msg_event_chat.id is '聊天记录唯一id(uuid格式)，默认空字符串';
comment on column t_msg_event_chat.parent is '父级聊天会话标识符，默认空字符串';
comment on column t_msg_event_chat.sender_call_sign is '发送方呼号标识(如"alpha-01")，默认空字符串';
comment on column t_msg_event_chat.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_chat.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_chat.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_chat.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_chat.is_deleted is '软删标志，默认false表示有效数据';


create unique index idx_unique_msg_event_chat_event_id on t_msg_event_chat(event_id);

-- 创建触发器
create trigger update_msg_event_chat_updated_at
before update on t_msg_event_chat
for each row
execute function update_timestamp();

-- 群组信息表
create table if not exists "t_msg_event_group" (
    event_id varchar(100) primary key,
    name varchar(100) not null,
    role varchar(100) not null default 'member',
    update_by int not null default 0,
    create_by int not null default 0,
    created_at bigint default extract(epoch from current_timestamp),
    updated_at bigint default extract(epoch from current_timestamp),
    is_deleted bool not null default 'f'
);
comment on table t_msg_event_group is '消息事件群组管理表-一级关联表';
comment on column t_msg_event_group.event_id is '事件唯一标识符(主键)';
comment on column t_msg_event_group.name is '群组名称(唯一标识，如"指挥中心-01")，不可为空';
comment on column t_msg_event_group.role is '成员角色(admin=管理员/member=普通成员)，默认member';
comment on column t_msg_event_group.update_by is '更新者id，默认0表示系统操作';
comment on column t_msg_event_group.create_by is '创建者id，默认0表示系统创建';
comment on column t_msg_event_group.created_at is '创建时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_group.updated_at is '更新时间(unix时间戳，秒级)，默认当前时间';
comment on column t_msg_event_group.is_deleted is '软删标志，默认false表示有效数据';


create unique index idx_unique_msg_event_group_event_id on t_msg_event_group(event_id);

-- 创建触发器
create trigger update_msg_event_group_updated_at
before update on t_msg_event_group
for each row
execute function update_timestamp();