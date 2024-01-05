#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += WORDS
WORDS_MK_SUMMARY := go-corelibs/words
WORDS_MK_VERSION := v1.0.0

include CoreLibs.mk
