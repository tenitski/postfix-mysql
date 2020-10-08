FROM ubuntu:focal

EXPOSE 25

RUN apt-get update && \
  echo "postfix postfix/mailname string example.com" | debconf-set-selections && \
  echo "postfix postfix/main_mailer_type string 'Internet Site'" | debconf-set-selections && \
  apt-get install -y \
    postfix \
    postfix-mysql \
    libsasl2-modules \
    libsasl2-modules-sql \
    sasl2-bin

RUN update-rc.d -f postfix remove

RUN postconf -e syslog_name=example-smtp
RUN postconf -e mynetworks=0.0.0.0/0

RUN postconf -e "maillog_file = /dev/stdout"

RUN cp /etc/host.conf /etc/hosts /etc/nsswitch.conf /etc/resolv.conf /etc/services /var/spool/postfix/etc

# See http://www.postfix.org/SASL_README.html

# Enable SASL
COPY smtpd.conf /etc/postfix/sasl/smtpd.conf
RUN postconf -e 'smtpd_sasl_path = smtpd'
RUN postconf -e 'smtpd_sasl_type = cyrus'
RUN postconf -e 'smtpd_sasl_auth_enable = yes'

# Enable sender address check
COPY smtpd_sender_login_maps.cf /etc/postfix/smtpd_sender_login_maps.cf
RUN postconf -e 'smtpd_sender_login_maps = mysql:/etc/postfix/smtpd_sender_login_maps.cf'
RUN postconf -e 'smtpd_relay_restrictions = reject_sender_login_mismatch permit_sasl_authenticated defer_unauth_destination'


CMD ["postfix", "start-fg"]
