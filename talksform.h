#ifndef TALKSFORM_H
#define TALKSFORM_H

#include <QWidget>

class TalksForm : public QWidget
{
    Q_OBJECT
public:
    explicit TalksForm(int talks, QWidget *parent = nullptr);
private:
    int _talks;

signals:

public slots:
};

#endif // TALKSFORM_H
