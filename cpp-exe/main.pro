#! /usr/bin/env pro
#QT = core
CONFIG -= QT

CONFIG += c++17 cmdline

SOURCES += \
        main.cpp

DESTDIR = $$PWD

INCLUDEPATH += $$(HOME)/common/include

LIBS += $$PWD/exportgo.a

#HEADERS += 
